package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/revel/revel"
	"github.com/satori/go.uuid"

	"golang.org/x/crypto/bcrypt"
)

var (
	USER_NEWBIE  = 0
	USER_AUTHOR  = 1
	USER_DOOMED  = 2
	USER_DELETED = 3
	USER_MOD     = 4
)

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

func (u UserModel) TableName() string {
	return "users"
}

func (u UserModel) String() string {
	user := "test"
	return fmt.Sprintf("User(%s) %s", u.Name, user)
}

func NewUser(Username string, Name string, Email string, Password string) UserModel {

	var hashedPassword, _ = bcrypt.GenerateFromPassword(
		[]byte(Password), bcrypt.DefaultCost)
	u4 := uuid.NewV4()
	user := UserModel{
		UserID:         u4.String(),
		Name:           Name,
		Slug:           revel.Slug(Username),
		Username:       Username,
		Email:          Email,
		Password:       Password,
		HashedPassword: string(hashedPassword),
		Type:           USER_NEWBIE,
	}

	return user
}
