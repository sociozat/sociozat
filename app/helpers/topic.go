package helpers

import (
	"sociozat/app/models"
	"github.com/jinzhu/gorm"
)

func TodaysTopics(db *gorm.DB, channels []uint) []models.TopicModel {
	topics := []models.TopicModel{}

	tx := db.Table("posts").
		Select("count(posts.id) as Post_Count, topics.name as Name, topics.slug as Slug").
		Joins("join topics on posts.topic_id = topics.id")

	if len(channels) > 0 {
		tx = tx.Joins("join topic_channels on posts.topic_id = topic_channels.topic_model_id").
			Where("topic_channels.channel_model_id IN(?)", channels)
	}

	tx.Limit(100).
		Group("topics.id").
		Order("topics.updated_at DESC").
		Find(&topics)

	return topics
}

func TrendingTopics(db *gorm.DB, channels []uint, startDate string) []models.TopicModel {
	topics := []models.TopicModel{}

	tx := db.Debug().Table("posts").
		Select("count(posts.id) as Post_Count, topics.name as Name, topics.slug as Slug").
        Where("posts.created_at >= ?", startDate).
		Joins("join topics on posts.topic_id = topics.id")

	if len(channels) > 0 {
		tx = tx.Joins("join topic_channels on posts.topic_id = topic_channels.topic_model_id").
			Where("topic_channels.channel_model_id IN(?)", channels)
	}

	tx.Limit(100).
		Group("topics.id").
		Order("topics.updated_at DESC").
		Find(&topics)

	return topics
}

