package repositories

import (
	"github.com/biezhi/gorm-paginator/pagination"
	"sozluk/app"
	"sozluk/app/models"
)

type TopicSearchParams struct {
	Slug      string
	StartDate string
	Page      int
	Limit     int
	OrderBy   []string
}

type TopicRepository struct{}

//Find topics by given criteria
func (t TopicRepository) Find(params TopicSearchParams) (models.TopicModel, *pagination.Paginator, error) {

	topic := models.TopicModel{}
	posts := []models.PostModel{}

	var err error

	if err := app.DB.Where("topics.slug = ?", params.Slug).Preload("Channels").Find(&topic).Error; err != nil {
		return topic, &pagination.Paginator{}, err
	}

	rows := app.DB.Table("posts").
		Where("posts.topic_id = ?", topic.ID).
		Where("posts.created_at >= ?", params.StartDate).
		Preload("User")

	paginator := pagination.Paging(&pagination.Param{
		DB:    rows,
		Page:  params.Page,
		Limit: params.Limit,
	}, &posts)

	return topic, paginator, err
}

//FindBySlug get topic info by slug
func (t TopicRepository) FindBySlug(slug string) (models.TopicModel, error) {
	var err error
	var topic = models.TopicModel{}
	if err := app.DB.Where("slug = ?", slug).First(&topic).Error; err != nil {
		return topic, err
	}

	return topic, err
}
