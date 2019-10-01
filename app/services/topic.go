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
	PostRepository    repositories.PostRepository
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

func (t TopicService) GetTopicbySlug(slug string) (models.TopicModel, error) {
	topic, err := t.TopicRepository.FindBySlug(slug)

	return topic, err
}

func (t TopicService) Reply(topic models.TopicModel, user *models.UserModel, content string) (*models.PostModel, error) {

	p := models.PostModel{
		Content: content,
		User:    user,
		Topic:   topic,
	}

	post, err := t.PostRepository.Create(&p)

	return post, err
}