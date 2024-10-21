package helpers

import (
	"sociozat/app/models"
	"github.com/jinzhu/gorm"
)

func TrendingTopics(db *gorm.DB, channels []uint, startDate string) []models.TopicModel {
	topics := []models.TopicModel{}

	tx := db.Table("topics").
		Select("count(p.id) as Post_Count, topics.name as Name, topics.slug as Slug, CASE WHEN SUM(p.likes - p.dislikes) = 0 THEN 0 ELSE LOG(GREATEST(SUM(p.likes - p.dislikes), 1)) + ((EXTRACT(EPOCH FROM NOW()) - EXTRACT(EPOCH FROM topics.created_at)) / 45000) END AS hotness_score").
        Where("p.created_at >= ?", startDate).
        Where("p.deleted_at is null").
		Joins("join posts  as p on topics.id = p.topic_id")

	if len(channels) > 0 {
		tx = tx.Joins("join topic_channels on topics.id = topic_channels.topic_model_id").
			Where("topic_channels.channel_model_id IN(?)", channels)
	}

	tx.Limit(100).
		Group("topics.id").
		Order("hotness_score DESC").
		Find(&topics)

	return topics
}

