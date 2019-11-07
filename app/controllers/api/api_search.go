package controllers

import (
	"fmt"
	"sociozat/app/models"
	"sociozat/app/repositories"

	"github.com/revel/revel"
)

type ApiSearch struct {
	*revel.Controller
	repositories.SearchRepository
}

func (c ApiSearch) Query(term string) revel.Result {

	adapter := revel.Config.StringDefault("search.adapter", "postgres")
	results := c.SearchRepository.SetAdapter(adapter).Query(term)

	action := make(map[string]string, 2)
	action["url"] = fmt.Sprintf("/search/%s", term)
	action["text"] = c.Message("search.all.results")
	response := models.SearchResponse{
		Results: results,
		Action:  action,
	}

	return c.RenderJSON(response)
}
