package services

import (
	"sociozat/app/models"
	"sociozat/app/repositories"

	"github.com/biezhi/gorm-paginator/pagination"
)

//ChannelService struct
type ChannelService struct {
	ChannelRepository repositories.ChannelRepository
}

//Search gets a channel list by search query
func (s ChannelService) Search(search string) ([]models.ChannelModel, error) {
	return s.ChannelRepository.Search(search)
}

//GetPostsByChannel gets a channel by slug or returns errror
func (s ChannelService) GetPostsByChannel(params models.SearchParams) (*pagination.Paginator, *models.ChannelModel, error) {
	return s.ChannelRepository.GetPostsByChannel(params)
}
