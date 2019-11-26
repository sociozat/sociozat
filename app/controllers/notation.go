package controllers

import (
	"github.com/gosimple/slug"
	"github.com/revel/revel"
	"net/url"
	"sociozat/app/services"
)

type Notation struct {
	App
	TopicService    services.TopicService
	SettingsService services.SettingsService
}

//View renders post by id
func (c Notation) View(name string) revel.Result {
	title, _ := url.QueryUnescape(name)
	topic, err := c.TopicService.GetTopicbySlug(slug.Make(title))
	if err == nil {
		return c.Redirect(Topic.View, topic.Slug)
	}

	return c.Render(title)
}
