package controllers

import (
    "sociozat/app"
	"sociozat/app/models"
	"sociozat/app/helpers"
	"sociozat/app/services"

	"github.com/gosimple/slug"
	"github.com/revel/revel"
	"github.com/mitchellh/mapstructure"
	"strconv"
	"time"
)

//App struct
type App struct {
	*revel.Controller
	UserService services.UserService
	PostService services.PostService
	TopicService services.TopicService
}

//Index renders home page
func (c App) Index() revel.Result {
	title, _ := revel.Config.String("app.name")

	posts, _ := c.PostService.GetHomePagePosts()

	return c.Render(title, posts)
}

//Index renders home page
func (c App) Trending() revel.Result {
	var ids []uint
    settings := models.SettingsModel{}
    sets, _ := c.Session.Get("settings")
    mapstructure.Decode(sets, &settings)
    for _, channel := range settings.TrendingChannels {
        ids = append(ids, channel.ID)
    }

    threshold, _ := strconv.Atoi(revel.Config.StringDefault("trending.threshold", "24"))
    currentTime := time.Now()
    startDate  := currentTime.Add(time.Duration(-threshold) * time.Hour).Format("2006-01-02 15:04:05") //set this as beginning

    var posts = helpers.TrendingTopics(app.DB, ids, startDate)

	return c.Render(posts)
}

//Index renders home page
func (c App) Today() revel.Result {
	title, _ := revel.Config.String("app.name")

	posts, _ := c.PostService.TodaysPosts()

	return c.Render(title, posts)
}

func (c App) connected() *models.UserModel {
	if c.ViewArgs["user"] != nil {
		return c.ViewArgs["user"].(*models.UserModel)
	}
	if username, ok := c.Session["user"]; ok {
		return c.GetUser(username.(string))
	}

	return nil
}

//GetUser gets user by username
func (c App) GetUser(username string) *models.UserModel {

	slug := slug.Make(username)
	user, err := c.UserService.GetBySlug(slug)

	if err == nil {
		return user
	}

	return nil
}
