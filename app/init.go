package app

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"sozluk/app/helpers"
	"sozluk/app/models"
	"strings"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/mitchellh/mapstructure"
	"github.com/revel/revel"
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
	dbName := revel.Config.StringDefault("db.name", "sozluk")
	dbUser := revel.Config.StringDefault("db.user", "root")
	dbPassword := revel.Config.StringDefault("db.password", "root")
	dbAddress := revel.Config.StringDefault("db.address", "127.0.0.1")
	dbPort := revel.Config.StringDefault("db.port", "3306")
	dbSslMode := revel.Config.StringDefault("db.sslmode", "disable")

	var connstring string
	if dbDriver == "mysql" {
		connstring = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True", dbUser, dbPassword, dbAddress, dbPort, dbName)
	} else if dbDriver == "postgres" {
		connstring = fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=%s", dbUser, dbPassword, dbAddress, dbName, dbSslMode)
	} else if dbDriver == "sqlite3" {
		connstring = dbAddress
	}

	var err error
	DB, err = gorm.Open(dbDriver, connstring)
	if err != nil {
		revel.AppLog.Info("DB Error", err)
	}
	revel.AppLog.Info("DB Connected")
	DB.AutoMigrate(&models.UserModel{})
	DB.AutoMigrate(&models.PostModel{})
	DB.AutoMigrate(&models.TopicModel{})
	DB.AutoMigrate(&models.ChannelModel{})
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
		Theme:          "default",
		TodaysChannels: nil,
		HeaderChannels: nil,
	}

	if sess, _ := c.Session.Get("settings"); sess == nil {
		var inInterface map[string]interface{}
		inrec, _ := json.Marshal(settings)
		json.Unmarshal(inrec, &inInterface)

		c.Session.Set("settings", inInterface)
	}

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

	revel.TemplateFuncs["todays"] = func(sets map[string]interface{}) []models.TopicModel {

		var ids []uint
		settings := models.SettingsModel{}
		mapstructure.Decode(sets, &settings)
		for _, channel := range settings.TodaysChannels {
			ids = append(ids, channel.ID)
		}

		return helpers.TodaysTopics(DB, ids)
	}

	revel.TemplateFuncs["format"] = func(str string) template.HTML {
		return template.HTML(strings.Replace(helpers.FormatContent(str), "\n", "<br>", -1))
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
