package services

import (
	"sozluk/app/models"
	"sozluk/app/repositories"
)

type ChannelService struct {
	ChannelRepository repositories.ChannelRepository
}

//List gets a channel list by searhc query
func (s ChannelService) List(search string) ([]models.ChannelModel, error) {
	return s.ChannelRepository.Find(search)
}
