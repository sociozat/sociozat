package search

import (
	"sociozat/app"
	"sociozat/app/models"
)

type MysqlSearch struct {
}

func (s MysqlSearch) Query(term string) []models.SiteSearchResult {
	//do query and set results
	results := []models.SiteSearchResult{}

	if err := app.DB.Table("searches").Where("WHERE MATCH (title) AGAINST (? IN BOOLEAN MODE)", term).Find(&results).Error; err != nil {
		return results
	}

	return results
}
