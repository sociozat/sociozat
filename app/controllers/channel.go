package controllers

import (
	"sozluk/app/services"

	"github.com/revel/revel"
)

type ChannelResponse struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type Message struct {
	Success bool              `json:"success"`
	Message string            `json:"message,omitempty"`
	Results []ChannelResponse `json:"results,omitempty"`
}

type Channel struct {
	App
	ChannelService services.ChannelService
}

//Json returns a json responce with channel collection
func (c Channel) Json() revel.Result {
	search := c.Params.Query.Get("s")
	revel.AppLog.Debug(search)
	if search == "" {
		return c.RenderJSON(Message{Success: false, Message: "Type Something"})
	}

	channels, _ := c.ChannelService.List(search)
	if len(channels) > 0 {
		//map channels as name-value
		response := []ChannelResponse{}
		for _, v := range channels {
			response = append(response, ChannelResponse{Name: v.Name, Value: v.Slug})
		}

		return c.RenderJSON(Message{Success: true, Results: response})
	}
	return c.RenderJSON(Message{Success: false, Message: "Nothing Found"})

}
