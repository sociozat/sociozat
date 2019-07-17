package repositories

import (
	"github.com/revel/revel"
	"sozluk/app"
	"sozluk/app/models"
)

type UserR struct{}

func (this UserR) Create(u models.UserM) (*models.UserM, error) {

	var err error

	if err := app.DB.Create(&u).Error; err != nil {
		revel.AppLog.Crit(err.Error())
	}

	return &u, err
}
