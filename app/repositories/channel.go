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

//First gets the channel by ID
func (c ChannelRepository) FindByID(cid int) (*models.ChannelModel, error) {

	channel := models.ChannelModel{}

	var err error
	if err := app.DB.Where(&channel).First(&channel, cid).Error; err != nil {
		return &channel, err
	}

	return &channel, err
}

//Create add new channel to db
func (c ChannelRepository) Create(channel models.ChannelModel) (*models.ChannelModel, error) {

	var err error
	if err := app.DB.Create(&channel).Error; err != nil {
		return &channel, err
	}

	return &channel, err
}
