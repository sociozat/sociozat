package helpers

import (
	"github.com/jinzhu/gorm"
	"sozluk/app/models"
	"time"
)

func TodaysTopics(db *gorm.DB) []models.TopicModel {
	topics := []models.TopicModel{}

	yesterday := time.Now().AddDate(0, 0, -1)

	db.Table("posts").
		Select("count(posts.id) as Post_Count, topics.name as Name, topics.slug as Slug").
		Joins("join topics on posts.topic_id = topics.id").
		Where("posts.created_at > ?", yesterday).
		Group("topics.id").
		Order("topics.updated_at DESC").
		Find(&topics)

	return topics
}
