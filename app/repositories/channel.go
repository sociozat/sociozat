package repositories

import (
	"errors"
	"sozluk/app"
	"sozluk/app/models"
)

type ChannelRepository struct{}

func (c ChannelRepository) List(model *models.ChannelModel) (*models.ChannelModel, error) {

	var err error
	channels := app.DB.Where(model).ScanRows()
	if channels.RecordNotFound() {
		err = errors.New("user not found")
	}

	return channels, err
}
