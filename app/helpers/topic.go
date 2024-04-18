package helpers

import (
	"sociozat/app/models"
	"github.com/jinzhu/gorm"
)

func TrendingTopics(db *gorm.DB, channels []uint, startDate string) []models.TopicModel {
	topics := []models.TopicModel{}

	tx := db.Table("topics").
		Select("count(p.id) as Post_Count, topics.name as Name, topics.slug as Slug, CASE  WHEN (SUM(p.likes) + SUM(p.dislikes)) = 0 THEN 0 ELSE (LN(SUM(p.likes) + SUM(p.dislikes)) / LN(10) * 2 + (EXTRACT(EPOCH FROM (NOW() - topics.created_at)) / 3600 / 24)) / (LN(SUM(p.likes) + SUM(p.dislikes)) / LN(10) + 1) END AS hotness_score").
        Where("p.created_at >= ?", startDate).
		Joins("join posts  as p on topics.id = p.topic_id")

	if len(channels) > 0 {
		tx = tx.Joins("join topic_channels on topics.id = topic_channels.topic_model_id").
			Where("topic_channels.channel_model_id IN(?)", channels)
	}

	tx.Limit(30).
		Group("topics.id").
		Order("hotness_score DESC").
		Find(&topics)

	return topics
}

