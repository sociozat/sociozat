package services

import (
	"sozluk/app/models"
	"sozluk/app/repositories"
)

type ChannelService struct {
	ChannelRepository repositories.ChannelRepository
}

func (s ChannelService) List(language string) []models.ChannelModel {

	model := models.ChannelModel{Language: language}

	return s.ChannelRepository.List(&model)
}
