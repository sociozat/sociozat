package services

import (
	"github.com/biezhi/gorm-paginator/pagination"
	"sozluk/app/models"
	"sozluk/app/repositories"
)

//TopicService struct
type TopicService struct {
	TopicRepository   repositories.TopicRepository
	TopicSearchParams repositories.TopicSearchParams
}

//FindBySlug find topic with posts by slug
func (t TopicService) FindBySlug(slug string, page int, limit int, date string) (models.TopicModel, *pagination.Paginator, error) {

	params := repositories.TopicSearchParams{
		Slug:      slug,
		StartDate: date,
		Page:      page,
		Limit:     limit,
		OrderBy:   []string{"id desc"},
	}

	topic, posts, err := t.TopicRepository.Find(params)

	if err != nil {
		return topic, posts, err
	}

	return topic, posts, err
}
