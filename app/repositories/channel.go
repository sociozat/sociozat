package repositories

import (
	"sozluk/app"
	"sozluk/app/models"

	"github.com/biezhi/gorm-paginator/pagination"
)

//ChannelRepository Struct
type ChannelRepository struct{}

//Search finds channels by search query
func (c ChannelRepository) Search(search string) ([]models.ChannelModel, error) {
	var channels []models.ChannelModel
	var err error
	if err := app.DB.Where("name LIKE ?", "%"+search+"%").Find(&channels).Error; err != nil {
		return channels, err
	}

	return channels, err
}

//FindByID gets the channel by ID
func (c ChannelRepository) FindByID(cid int) (*models.ChannelModel, error) {

	channel := models.ChannelModel{}

	var err error
	if err := app.DB.Where(&channel).First(&channel, cid).Error; err != nil {
		return &channel, err
	}

	return &channel, err
}

//GetPostsByChannel gets the channel by ID
func (c ChannelRepository) GetPostsByChannel(params models.SearchParams) (*pagination.Paginator, *models.ChannelModel, error) {

	channel := models.ChannelModel{}
	posts := []models.PostModel{}

	if err := app.DB.Where("slug = ?", params.Slug).First(&channel).Error; err != nil {
		return new(pagination.Paginator), &channel, err
	}

	var err error

	tx := app.DB.Debug().Table("posts").
		Joins("join topics on topics.id = posts.topic_id").
		Joins("join topic_channels on posts.topic_id = topic_channels.topic_model_id").
		Where("topic_channels.channel_model_id = ?", channel.ID).
		Preload("Topic")

	paginator := pagination.Paging(&pagination.Param{
		DB:      tx,
		Page:    params.Page,
		Limit:   params.Limit,
		OrderBy: params.OrderBy,
	}, &posts)

	if err := tx.Error; err != nil {
		return paginator, &channel, err
	}

	return paginator, &channel, err
}

//Create add new channel to db
func (c ChannelRepository) Create(channel models.ChannelModel) (*models.ChannelModel, error) {

	var err error
	if err := app.DB.Create(&channel).Error; err != nil {
		return &channel, err
	}

	return &channel, err
}
