package search

import "sociozat/app/models"

type MysqlSearch struct {
}

func (s MysqlSearch) Query(term string) []models.SiteSearchResult {
	//do query and set results
	tmp := make([]models.SiteSearchResult, 10)
	for i := range tmp {
		tmp[i].Title = "title coming from mysql"
		tmp[i].Slug = "url"
		tmp[i].Type = "post"
	}

	return tmp
}
