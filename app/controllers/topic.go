package controllers

import (
	"github.com/revel/revel"
	"sozluk/app/services"
	"strconv"
)

type Topic struct {
	App
	TopicService services.TopicService
}

//View renders post by id
func (c Topic) View(slug string) revel.Result {

	page, _ := strconv.Atoi(c.Params.Query.Get("page"))

	limit, _ := strconv.Atoi(c.Params.Query.Get("limit"))

	startDate := c.Params.Query.Get("start_date")
	if startDate == "" {
		startDate = "1970-01-01" //set this as beginning
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

	return c.Render(title, topic, posts)
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

	return c.Redirect(Post.View, post.ID)
}
