package search

import (
	"sociozat/app"
	"sociozat/app/models"
)

type PostgresSearch struct {
}

func (s PostgresSearch) Query(term string) []models.SiteSearchResult {

	results := []models.SiteSearchResult{}

	if err := app.DB.Debug().Table("searches").Where("title @@ to_tsquery(?)", term).Find(&results).Error; err != nil {
		return results
	}

	return results
}
