package controllers

import (
    "sociozat/app"
	"sociozat/app/models"
	"sociozat/app/helpers"
	"sociozat/app/services"
	"sociozat/app/websocket"

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


func (c App) PostsSocket(ws revel.ServerWebSocket) revel.Result {
	// Make sure the websocket is valid.
	if ws == nil {
		return nil
	}

	// Join the room.
	subscription := websocket.Subscribe()
	defer subscription.Cancel()

	// Send down the archive.
	for _, event := range subscription.Archive {
		if ws.MessageSendJSON(&event) != nil {
			// They disconnected
			return nil
		}
	}

	// In order to select between websocket messages and subscription events, we
	// need to stuff websocket events into a channel.
	newMessages := make(chan string)
	go func() {
		var msg string
		for {
			err := ws.MessageReceiveJSON(&msg)
			if err != nil {
				close(newMessages)
				return
			}
			newMessages <- msg
		}
	}()

	// Now listen for new events from either the websocket
	for {
		select {
		case event := <-subscription.New:
			if ws.MessageSendJSON(&event) != nil {
				// They disconnected.
				return nil
			}
		case msg, ok := <-newMessages:
			// If the channel is closed, they disconnected.
			if !ok {
				return nil
			}

			// Otherwise, say something.
			websocket.Publish("message", msg)
		}
	}

	return nil
}