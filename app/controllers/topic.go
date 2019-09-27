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
