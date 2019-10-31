package controllers

import (
	"fmt"
	"sociozat/app/models"
	"sociozat/app/services"
	"strconv"

	"github.com/revel/revel"
)

type ChannelResponse struct {
	Name  string `json:"name"`
	Value uint   `json:"value"`
}

type Message struct {
	Success bool              `json:"success"`
	Message string            `json:"message,omitempty"`
	Results []ChannelResponse `json:"results,omitempty"`
}

type Channel struct {
	App
	ChannelService  services.ChannelService
	SettingsService services.SettingsService
}

//Json returns a json responce with channel collection
func (c Channel) Json() revel.Result {
	search := c.Params.Query.Get("s")
	if search == "" {
		return c.RenderJSON(Message{Success: false, Message: "Type Something"})
	}

	channels, _ := c.ChannelService.Search(search)
	if len(channels) > 0 {
		//map channels as name-value
		response := []ChannelResponse{}
		for _, v := range channels {
			response = append(response, ChannelResponse{Name: v.Name, Value: v.ID})
		}

		return c.RenderJSON(Message{Success: true, Results: response})
	}
	return c.RenderJSON(Message{Success: false, Message: "Nothing Found"})

}

//View channel detail page with posts
func (c Channel) View(slug string) revel.Result {
	page, _ := strconv.Atoi(c.Params.Query.Get("page"))

	sets, err := c.Session.Get("settings")

	settings, _ := c.SettingsService.MapSettings(sets)
	params := models.SearchParams{
		Slug:    slug,
		Page:    page,
		Limit:   settings.PostPerPage,
		OrderBy: []string{"posts.id DESC"},
	}

	posts, channel, err := c.ChannelService.GetPostsByChannel(params)

	if err != nil {
		c.Flash.Error(err.Error())
		return c.Redirect(App.Index)
	}

	var pagination = make(map[int]string)
	for i := 1; i <= posts.TotalPage; i++ {
		c.Params.Query.Set("page", strconv.Itoa(i))
		pageValue := fmt.Sprintf("/c/%s?%s", c.Params.Route.Get("slug"), c.Params.Query.Encode())
		pagination[i] = pageValue
	}

	return c.Render(posts, channel, pagination)
}
