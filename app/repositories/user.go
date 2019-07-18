package repositories

import (
	// "github.com/revel/revel"
	"sozluk/app"
	"sozluk/app/models"
)

type UserRepository struct{}

func (this UserRepository) Create(u models.UserModel) (*models.UserModel, error) {

	var err error

	if err := app.DB.Create(&u).Error; err != nil {
		return &u, err
	}

	return &u, err
}
