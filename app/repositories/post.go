package repositories

import (
	"sozluk/app"
	"sozluk/app/models"
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

func (c PostRepository) FirstByID(id int) (*models.PostModel, error) {
	post := models.PostModel{}

	var err error
	if err := app.DB.Where(&post).Preload("Topic").Preload("User").First(&post, id).Error; err != nil {
		return &post, err
	}

	return &post, err
}
