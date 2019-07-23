package controllers

import (
	"github.com/revel/revel"
	"sozluk/app"
	"sozluk/app/services"
)

type Channel struct {
	ChannelService services.ChannelService
}

func (c Channel) JsonList() revel.Result {

	channels := c.ChannelService.List(app.DefaultLocale)
	return revel.RenderJSONResult()
}
