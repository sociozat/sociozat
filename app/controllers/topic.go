package controllers

import (
	"fmt"
	"sociozat/app/services"
	"sociozat/app/websocket"
	"strconv"
	"time"
	"strings"

	"github.com/revel/revel"
)

type Topic struct {
	App
	TopicService    services.TopicService
	SettingsService services.SettingsService
}

//View renders post by id
func (c Topic) View(slug string) revel.Result {

	page, _ := strconv.Atoi(c.Params.Query.Get("page"))

	sets, _ := c.Session.Get("settings")

	settings, _ := c.SettingsService.MapSettings(sets)
	limit := 10
	if(settings.PostPerPage > 0){
	    limit = settings.PostPerPage
	}

	a := c.Params.Query.Get("a")

    startDate  := "1970-01-01"
    currentTime := time.Now()

	if a == "trending" {
        threshold, _ := strconv.Atoi(revel.Config.StringDefault("trending.threshold", "24"))
		startDate  = currentTime.Add(time.Duration(-threshold) * time.Hour).Format("2006-01-02 15:04:05") //set this as beginning
	}

	if a == "today" {
		startDate  = currentTime.Add(time.Duration(-24) * time.Hour).Format("2006-01-02 15:04:05")
	}

	topic, posts, err := c.TopicService.FindBySlug(slug, page, limit, startDate)
	if err != nil {
		c.FlashParams()
		return c.Redirect(App.Index)
	}

	var title = c.Message("topic.title", topic.Name)


	if page > 0 {
		title = c.Message("topic.title.with.page", topic.Name, page)
	}

    //add total post
    previousPostCount := 0
    previousPostsPage := 0
    if a == "trending" ||  a == "today" {
        previousPostCount  = c.TopicService.PostCountUntil(topic, startDate)
        previousPostsPage = int(previousPostCount) / int(limit)
    }

	//set pages
	c.Params.Query = c.Request.URL.Query()

	var pagination = make(map[int]string)
	for i := 1; i <= posts.TotalPage; i++ {
		c.Params.Query.Set("page", strconv.Itoa(i))

		pageValue := fmt.Sprintf("/t/%s?%s", c.Params.Route.Get("slug"), c.Params.Query.Encode())
		pagination[i] = pageValue
	}

    canonical := fmt.Sprintf("%s/t/%s",  revel.Config.StringDefault("app.url", ""), c.Params.Route.Get("slug"))
	return c.Render(title, topic, canonical, posts, pagination, previousPostCount, previousPostsPage)
}

//Reply topic with POST method
func (c Topic) Reply(slug string) revel.Result {
	u := c.connected()

	t, err := c.TopicService.GetTopicbySlug(slug)

	if u == nil {
		c.Flash.Error(c.Message("auth.login.required"))
		return c.Redirect(Topic.View, t.ID)
	}

	if err != nil {
		c.Flash.Error(err.Error())
		return c.Redirect(Topic.View, t.ID)
	}

	content := c.Params.Form.Get("content")
	post, err := c.TopicService.Reply(t, u, content)

	if err != nil {
		c.Flash.Error(err.Error())
		return c.Redirect(Topic.View, t.ID)
	}

    //publish via websocket
    var list []string
    for _, c := range t.Channels {
        list = append(list, strconv.FormatUint(uint64(c.ID), 10))
    }
	websocket.Publish("channels", strings.Join(list, ","))

	c.Flash.Success(c.Message("topic.create.success.message"))

	return c.Redirect(Post.View, post.ID)
}
