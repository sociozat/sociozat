package services

import (
	"github.com/revel/revel"
	"regexp"
	"sozluk/app"
	"sozluk/app/models"
	"sozluk/app/repositories"
)

type UserService struct {
	UserRepository repositories.UserRepository
}

var userRegex = regexp.MustCompile("^\\w*$")

func (this UserService) Create(m models.UserModel, rv *revel.Validation) (*models.UserModel, map[string]*revel.ValidationError, error) {

	v := this.Validate(m, rv)

	var err error

	//save to db
	if v == nil {
		newUser := models.NewUser(m.Username, m.Username, m.Email, m.Password)

		u, err := this.UserRepository.Create(newUser)

		if err != nil {
			return &m, v, err
		}

		return u, v, err
	}

	return &m, v, err
}

func (this UserService) Validate(m models.UserModel, rv *revel.Validation) map[string]*revel.ValidationError {
	rv.Check(m.Username,
		revel.Required{},
		revel.MaxSize{15},
		revel.MinSize{4},
		revel.Match{userRegex},
	).Message(app.Trans("user.create.validation.username"))

	rv.Check(m.Password,
		revel.Required{},
		revel.MaxSize{15},
		revel.MinSize{5},
	).Message(app.Trans("user.create.validation.password"))

	if rv.HasErrors() {
		rv.Keep()
		return rv.ErrorMap()
	}

	return nil
}
