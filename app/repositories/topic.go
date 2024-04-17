package repositories

import (
	"github.com/biezhi/gorm-paginator/pagination"
	"sociozat/app"
	"sociozat/app/models"
	"time"
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
		Order("posts.id ASC").
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
	if err := app.DB.Where("slug = ?", slug).Preload("Channels").First(&topic).Error; err != nil {
		return topic, err
	}

	return topic, err
}

type TotalCount struct {
  Total int
}

func (t TopicRepository) CountPostsUntil(ID uint, date string) int {

    var result TotalCount

    app.DB.Table("posts").Select("count(id) as Total").
        Where("posts.topic_id = ?", ID).
        Where("posts.created_at <= ?", date).
        Scan(&result)

    return result.Total

}

func (c TopicRepository) Todays(limit int) ([]models.TopicModel, error) {
    topics := []models.TopicModel{}

    currentTime := time.Now()
    startDate  := currentTime.Add(time.Duration(-24) * time.Hour).Format("2006-01-02 15:04:05")

    tx := app.DB.Table("posts").
        Select("count(posts.id) as Post_Count, topics.name as Name, topics.slug as Slug").
        Joins("join topics on posts.topic_id = topics.id").
        Where("posts.created_at >= ?", startDate).
        Group("topics.id").
        Order("topics.updated_at ASC").
        Limit(limit)


    if err := tx.Find(&topics).Error; err != nil {
		return topics, err
	}

    return topics, nil

}