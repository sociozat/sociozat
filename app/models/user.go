package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/satori/go.uuid"
	"regexp"

	"github.com/revel/revel"
	"golang.org/x/crypto/bcrypt"
)

var (
	USER_NEWBIE  = 0
	USER_AUTHOR  = 1
	USER_DOOMED  = 2
	USER_DELETED = 3
	USER_MOD     = 4
)

type UserM struct {
	gorm.Model
	UserID         string `gorm:"type:char(36);primary_key"`
	Name           string
	Username       string `gorm:"unique;not null"`
	Email          string `gorm:"unique;not null"`
	Password       string `gorm:"-"`
	HashedPassword string `gorm:"type:varchar(255)"`
	Type           int
}

func (u UserM) String() string {
	user := "test"
	return fmt.Sprintf("User(%s) %s", u.Name, user)
}

var userRegex = regexp.MustCompile("^\\w*$")

func (user *UserM) Validate(v *revel.Validation) {
	v.Check(user.Username,
		revel.Required{},
		revel.MaxSize{15},
		revel.MinSize{4},
		revel.Match{userRegex},
	)

	ValidatePassword(v, user.Password).
		Key("user.Password")

}

func ValidatePassword(v *revel.Validation, password string) *revel.ValidationResult {
	return v.Check(password,
		revel.Required{},
		revel.MaxSize{15},
		revel.MinSize{5},
	)
}

func NewUser(Username string, Name string, Email string, Password string) UserM {

	var hashedPassword, _ = bcrypt.GenerateFromPassword(
		[]byte(Password), bcrypt.DefaultCost)
	u4 := uuid.NewV4()
	user := UserM{
		UserID:         u4.String(),
		Name:           Name,
		Username:       Username,
		Email:          Email,
		Password:       Password,
		HashedPassword: string(hashedPassword),
		Type:           USER_NEWBIE,
	}

	return user
}
