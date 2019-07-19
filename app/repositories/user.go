package repositories

import (
	// "github.com/revel/revel"
	"errors"
	"sozluk/app"
	"sozluk/app/models"
)

//UserRepository struct
type UserRepository struct{}

//Create add new user to db
func (c UserRepository) Create(u models.UserModel) (models.UserModel, error) {

	var err error

	if err := app.DB.Create(&u).Error; err != nil {
		return u, err
	}

	return u, err
}

//GetUserBySlug get user from database
func (c UserRepository) GetUserById(Id uint) (u *models.UserModel, err error) {
	user := &models.UserModel{}
	record := app.DB.Where("id=?", Id).Find(&user)
	if record.RecordNotFound() {
		err = errors.New("user not found")
	}
	return user, err
}

//GetUserBySlug get user from database
func (c UserRepository) GetUserBySlug(slug interface{}) (u *models.UserModel, err error) {
	user := &models.UserModel{}
	record := app.DB.Where("slug=?", slug).Find(&user)
	if record.RecordNotFound() {
		err = errors.New("user not found")
	}
	return user, err
}
