package repositories

import (
	"log"
	"sociozat/app/models"
	"sociozat/app/repositories/search"
)

type SearchInterface interface {
	Query(term string) []models.SiteSearchResult
}

type SearchRepository struct {
	Adapter string
}

//SetAdapter sets database type for search
func (s *SearchRepository) SetAdapter(adapter string) *SearchRepository {
	s.Adapter = adapter
	return s
}

//GetAdapterType creates database adapter for search
func (s SearchRepository) GetAdapterType(adapter string) SearchInterface {
	switch s.Adapter {
	case "postgres":
		return search.PostgresSearch{}
	case "mysql":
		return search.MysqlSearch{}
	default:
		log.Printf("adapter undefined, using postgres search as default ")
		return search.PostgresSearch{}
	}

}

//Query makes the searchover database adapter
func (s SearchRepository) Query(term string) []models.SiteSearchResult {
	adapter := s.GetAdapterType(s.Adapter)
	return adapter.Query(term)
}
