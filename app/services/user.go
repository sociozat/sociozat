package services

import (
	"github.com/revel/revel"
	"regexp"
	"sozluk/app/models"
	"sozluk/app/repositories"
)

type UserService struct {
	Validation     *revel.Validation
	UserRepository repositories.UserRepository
}

var userRegex = regexp.MustCompile("^\\w*$")

func (this UserService) Create(m models.UserModel, rv *revel.Validation) (*models.UserModel, map[string]*revel.ValidationError, error) {

	v := this.Validate(m, rv)

	var err error
	newUser := models.NewUser(m.Username, m.Username, m.Email, m.Password)

	//save to db
	u, err := this.UserRepository.Create(newUser)

	return u, v, err
}

func (this UserService) Validate(m models.UserModel, rv *revel.Validation) map[string]*revel.ValidationError {
	rv.Check(m.Username,
		revel.Required{},
		revel.MaxSize{15},
		revel.MinSize{4},
		revel.Match{userRegex},
	)

	rv.Check(m.Password,
		revel.Required{},
		revel.MaxSize{15},
		revel.MinSize{5},
	)

	if rv.HasErrors() {
		rv.Keep()
		return rv.ErrorMap()
	}

	return nil
}
