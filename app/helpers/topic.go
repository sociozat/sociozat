package helpers

import (
	"sozluk/app/models"
	"time"

	"github.com/jinzhu/gorm"
)

func TodaysTopics(db *gorm.DB, channels []uint) []models.TopicModel {
	topics := []models.TopicModel{}

	yesterday := time.Now().AddDate(0, 0, -1)

	tx := db.Table("posts").
		Select("count(posts.id) as Post_Count, topics.name as Name, topics.slug as Slug").
		Joins("join topics on posts.topic_id = topics.id")

	if len(channels) > 0 {
		tx = tx.Joins("join topic_channels on posts.topic_id = topic_channels.topic_model_id").
			Where("topic_channels.channel_model_id IN(?)", channels)
	}

	tx.Where("posts.created_at > ?", yesterday).
		Group("topics.id").
		Order("topics.updated_at DESC").
		Find(&topics)

	return topics
}
