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
