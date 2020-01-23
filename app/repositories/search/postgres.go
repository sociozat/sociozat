package search

import (
	"sociozat/app"
	"sociozat/app/models"
)

type PostgresSearch struct {
}

func (s PostgresSearch) Query(term string) []models.SiteSearchResult {

	results := []models.SiteSearchResult{}

	if err := app.DB.Table("searches").Where("title @@ plainto_tsquery(?)", term).Find(&results).Error; err != nil {
		return results
	}

	return results
}
