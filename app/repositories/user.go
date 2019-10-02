package repositories

import (
	"github.com/biezhi/gorm-paginator/pagination"
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

func (c UserRepository) GetUserInfo(params models.SearchParams) (models.UserModel, *pagination.Paginator, error) {

	user := models.UserModel{}
	posts := []models.PostModel{}

	var err error

	if err := app.DB.Where("users.slug = ?", params.Slug).Find(&user).Error; err != nil {
		return user, &pagination.Paginator{}, err
	}

	rows := app.DB.Table("posts").
		Where("posts.user_id = ?", user.ID).
		Preload("Topic")

	paginator := pagination.Paging(&pagination.Param{
		DB:    rows,
		Page:  params.Page,
		Limit: params.Limit,
	}, &posts)

	return user, paginator, err
}
