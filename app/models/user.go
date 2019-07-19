package models

import (
	"github.com/gosimple/slug"
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
)

const (
	USER_NEWBIE  = 0
	USER_AUTHOR  = 1
	USER_DOOMED  = 2
	USER_DELETED = 3
	USER_MOD     = 4
)

//UserModel struct
type UserModel struct {
	gorm.Model
	UserID         string `gorm:"type:char(36);primary_key"`
	Slug           string `gorm:"index:user_slug"`
	Name           string
	Username       string `gorm:"unique;not null"`
	Email          string `gorm:"unique;not null"`
	Password       string `gorm:"-"`
	HashedPassword string `gorm:"type:varchar(255)"`
	Type           int
}

//TableName sets table name on db
func (u UserModel) TableName() string {
	return "users"
}

//NewUser returns new user instance
func NewUser(Username string, Name string, Email string, Password string) UserModel {

	var hashedPassword, _ = bcrypt.GenerateFromPassword(
		[]byte(Password), bcrypt.DefaultCost)
	u4 := uuid.NewV4()
	user := UserModel{
		UserID:         u4.String(),
		Name:           Name,
		Slug:           slug.Make(Username),
		Username:       Username,
		Email:          Email,
		Password:       Password,
		HashedPassword: string(hashedPassword),
		Type:           USER_NEWBIE,
	}

	return user
}
