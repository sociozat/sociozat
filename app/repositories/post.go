package repositories

import (
	"sociozat/app"
	"sociozat/app/models"
)

type PostRepository struct{}

//Create add new post to db
func (c PostRepository) Create(p *models.PostModel) (*models.PostModel, error) {

	var err error
	if err := app.DB.Create(&p).Error; err != nil {
		return p, err
	}

	return p, err
}

//Update post by id
func (c PostRepository) Update(p *models.PostModel) (*models.PostModel, error) {

	var err error
	if err := app.DB.Save(&p).Error; err != nil {
		return p, err
	}

	return p, err
}

func (c PostRepository) FindByID(id int) (*models.PostModel, error) {
	post := models.PostModel{}

	var err error
	if err := app.DB.Where(&post).Preload("User").First(&post, id).Error; err != nil {
		return &post, err
	}

	//add topic channels
	topic := models.TopicModel{}
	app.DB.Where("id = ?", post.TopicID).Preload("Channels").Find(&topic)
	post.Topic = topic

	return &post, err
}
