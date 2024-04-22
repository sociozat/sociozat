package app

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"sociozat/app/helpers"
	"sociozat/app/models"
	"strings"
	"strconv"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/mitchellh/mapstructure"
	"github.com/revel/revel"
	"github.com/grokify/html-strip-tags-go"
)

var (
	// AppVersion revel app version (ldflags)
	AppVersion string

	// BuildTime revel app build-time (ldflags)
	BuildTime string
)

var DB *gorm.DB

func InitDB() {

	dbDriver := revel.Config.StringDefault("db.driver", "mysql")
	dbName := revel.Config.StringDefault("db.name", "sociozat")
	dbUser := revel.Config.StringDefault("db.user", "root")
	dbPassword := revel.Config.StringDefault("db.password", "root")
	dbAddress := revel.Config.StringDefault("db.address", "127.0.0.1")
	dbPort := revel.Config.StringDefault("db.port", "3306")
	dbSslMode := revel.Config.StringDefault("db.sslmode", "disable")
	dbDebug := revel.Config.BoolDefault("db.debug", false)
	dbOptions := revel.Config.StringDefault("db.options", "-")

	var connstring string
	if dbDriver == "mysql" {
		connstring = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True", dbUser, dbPassword, dbAddress, dbPort, dbName)
	} else if dbDriver == "postgres" {
		connstring = fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s&options=%s", dbUser, dbPassword, dbAddress, dbPort, dbName, dbSslMode, dbOptions)
	} else if dbDriver == "sqlite3" {
		connstring = dbAddress
	}

	var err error
	DB, err = gorm.Open(dbDriver, connstring)
	if err != nil {
		revel.AppLog.Info("DB Error", err)
	}
	revel.AppLog.Info("DB Connected")
	DB.LogMode(dbDebug)
	DB.DB().SetMaxIdleConns(10)
	DB.DB().SetMaxOpenConns(100)
	DB.DB().SetConnMaxLifetime(time.Hour)

	DB.AutoMigrate(&models.UserModel{})
	DB.AutoMigrate(&models.PostModel{})
	DB.AutoMigrate(&models.TopicModel{})
	DB.AutoMigrate(&models.ChannelModel{})
	DB.AutoMigrate(&models.InvitationModel{})
	DB.AutoMigrate(&models.PostActionModel{})
}

var DefaultLocale string
var GetDefaultLocaleFilter = func(c *revel.Controller, fc []revel.Filter) {
	DefaultLocale = c.Request.Locale
	fc[0](c, fc[1:]) // Execute the next filter stage.
}

//for using outside of controller
func Trans(msg string, args ...interface{}) string {
	if len(args) > 0 {
		return revel.Message(DefaultLocale, msg, args)
	}
	return revel.Message(DefaultLocale, msg)

}

var CurrentUser *models.UserModel
var GetCurrenUser = func(c *revel.Controller, fc []revel.Filter) {

	u := models.UserModel{}

	if err := json.Unmarshal([]byte(c.Session["user"].(string)), &u); err == nil {
		CurrentUser = &u
	}

	fc[0](c, fc[1:]) // Execute the next filter stage.
}

//SetDefaultSettings if there is no settings in session.
func SetDefaultSettings(c *revel.Controller) revel.Result {
	//set default settings
	settings := models.SettingsModel{
		PostPerPage:    10,
		TopicPerPage:   10,
		Theme:          "default",
		TrendingChannels: nil,
		HeaderChannels: nil,
	}

	if sess, _ := c.Session.Get("settings"); sess == nil {
		var inInterface map[string]interface{}
		inrec, _ := json.Marshal(settings)
		json.Unmarshal(inrec, &inInterface)

		c.Session.Set("settings", inInterface)
	}

    c.Session.Set("isHtmx", c.Request.Header.Get("HX-Request"))
	return nil
}

func init() {
	// Filters is the default set of global filters.
	revel.Filters = []revel.Filter{
		revel.PanicFilter,             // Recover from panics and display an error page instead.
		revel.RouterFilter,            // Use the routing table to select the right Action
		revel.FilterConfiguringFilter, // A hook for adding or removing per-Action filters.
		revel.ParamsFilter,            // Parse parameters into Controller.Params.
		revel.SessionFilter,           // Restore and write the session cookie.
		revel.FlashFilter,             // Restore and write the flash cookie.
		revel.ValidationFilter,        // Restore kept validation errors and save new ones from cookie.
		revel.I18nFilter,              // Resolve the requested language
		GetDefaultLocaleFilter,        // Get Default Locale
		HeaderFilter,                  // Add some security based headers
		revel.InterceptorFilter,       // Run interceptors around the action.
		revel.CompressFilter,          // Compress the result.
		revel.BeforeAfterFilter,       // Call the before and after filter functions
		revel.ActionInvoker,           // Invoke the action.
	}

	revel.InterceptMethod(SetDefaultSettings, revel.BEFORE)

	// Register startup functions with OnAppStart
	// revel.DevMode and revel.RunMode only work inside of OnAppStart. See Example Startup Script
	// ( order dependent )
	// revel.OnAppStart(ExampleStartupScript)
	// revel.OnAppStart(InitDB)
	// revel.OnAppStart(FillCache)

	//read config function for templates
	revel.TemplateFuncs["config"] = func(a string, b string) string {
		return revel.Config.StringDefault(a, b)
	}

	revel.TemplateFuncs["user"] = func() *models.UserModel {
		return CurrentUser
	}

	revel.TemplateFuncs["trending"] = func(sets map[string]interface{}) []models.TopicModel {

		var ids []uint
		settings := models.SettingsModel{}
		mapstructure.Decode(sets, &settings)
		for _, channel := range settings.TrendingChannels {
			ids = append(ids, channel.ID)
		}

        threshold, _ := strconv.Atoi(revel.Config.StringDefault("trending.threshold", "24"))
        currentTime := time.Now()
		startDate  := currentTime.Add(time.Duration(-threshold) * time.Hour).Format("2006-01-02 15:04:05") //set this as beginning

		return helpers.TrendingTopics(DB, ids, startDate)
	}

	revel.TemplateFuncs["format"] = func(str string) template.HTML {
		var content string
		content = strip.StripTags(str)
		return template.HTML(strings.Replace(helpers.FormatContent(content), "\n", "<br>", -1))
	}

	revel.TemplateFuncs["marshal"] = func(v interface{}) string {
		a, _ := JSONMarshal(v)
		return string(a)
	}

	revel.OnAppStart(InitDB)
}

// HeaderFilter adds common security headers
// There is a full implementation of a CSRF filter in
// https://github.com/revel/modules/tree/master/csrf
var HeaderFilter = func(c *revel.Controller, fc []revel.Filter) {
	c.Response.Out.Header().Add("X-Frame-Options", "SAMEORIGIN")
	c.Response.Out.Header().Add("X-XSS-Protection", "1; mode=block")
	c.Response.Out.Header().Add("X-Content-Type-Options", "nosniff")
	c.Response.Out.Header().Add("Referrer-Policy", "strict-origin-when-cross-origin")

	fc[0](c, fc[1:]) // Execute the next filter stage.
}

func JSONMarshal(t interface{}) ([]byte, error) {
	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(true)
	err := encoder.Encode(t)
	return buffer.Bytes(), err
}

//func ExampleStartupScript() {
//	// revel.DevMod and revel.RunMode work here
//	// Use this script to check for dev mode and set dev/prod startup scripts here!
//	if revel.DevMode == true {
//		// Dev mode
//	}
//}
