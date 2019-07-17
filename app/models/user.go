package models

import (
	"fmt"
	"regexp"

	"github.com/revel/revel"
	"github.com/twinj/uuid"
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
	ID             uuid.UUID
	Name           string
	Username       string
	Email          string
	Password       string
	HashedPassword []byte
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

	v.Check(user.Name,
		revel.Required{},
		revel.MaxSize{100},
	)
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

	return UserM{
		uuid.NewV4(),
		Name,
		Username,
		Email,
		Password,
		hashedPassword,
		USER_NEWBIE,
	}
}
