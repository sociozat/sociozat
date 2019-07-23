package repositories

import (
	"sozluk/app"
	"sozluk/app/models"
)

type ChannelRepository struct{}

//Find finds channels by search query
func (c ChannelRepository) Find(search string) ([]models.ChannelModel, error) {
	var channels []models.ChannelModel
	var err error
	if err := app.DB.Where("name LIKE ?", "%"+search+"%").Find(&channels).Error; err != nil {
		return channels, err
	}

	return channels, err
}
