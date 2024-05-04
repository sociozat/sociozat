package models

import (
	"encoding/json"

	"github.com/gosimple/slug"
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
)

const (
	UserNewbie  = 1
	UserAuthor  = 2
	UserDoomed  = 3
	UserDeleted = 4
	UserMod     = 5
)

//UserModel struct
type UserModel struct {
	gorm.Model
	UserID         string `gorm:"type:char(36)"`
	Slug           string `gorm:"index:user_slug"`
	Name           string
	Username       string `gorm:"unique;not null"`
	Email          string `gorm:"unique;not null"`
	Password       string `gorm:"-"`
	HashedPassword string `gorm:"type:varchar(255)"`
	Settings       string `gorm:"type:text`
	Type           int
	EmailVerified  bool   `gorm:"type:bool`
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

	settings, _ := json.Marshal(&SettingsModel{})
	user := UserModel{
		UserID:         u4.String(),
		Name:           Name,
		Slug:           slug.Make(Username),
		Username:       Username,
		Email:          Email,
		Password:       Password,
		HashedPassword: string(hashedPassword),
		Settings:       string(settings),
		Type:           UserNewbie,
	}

	return user
}
