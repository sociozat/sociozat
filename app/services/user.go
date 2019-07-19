package services

import (
	"regexp"
	"sozluk/app"
	"sozluk/app/models"
	"sozluk/app/repositories"

	"github.com/revel/revel"
)

//UserService struct
type UserService struct {
	UserRepository repositories.UserRepository
}

var userRegex = regexp.MustCompile("^\\w*$")

//Create create new user or returns validation error
func (c UserService) Create(m models.UserModel, rv *revel.Validation) (models.UserModel, map[string]*revel.ValidationError, error) {

	v := c.Validate(m, rv)

	var err error

	//save to db
	if v == nil {
		newUser := models.NewUser(m.Username, m.Username, m.Email, m.Password)
		u, err := c.UserRepository.Create(newUser)
		if err != nil {
			return m, v, err
		}
		return u, v, err
	}
	return m, v, err
}

//GetBySlug returns err or user instance from database
func (c UserService) GetBySlug(Slug string) (u *models.UserModel, err error) {
	return c.UserRepository.GetUserBySlug(Slug)
}

//Validate validates user model form
func (c UserService) Validate(m models.UserModel, rv *revel.Validation) map[string]*revel.ValidationError {
	rv.Check(m.Username,
		revel.Required{},
		revel.MaxSize{25},
		revel.MinSize{4},
	).Message(app.Trans("user.create.validation.username"))

	rv.Check(m.Password,
		revel.Required{},
		revel.MaxSize{50},
		revel.MinSize{5},
	).Message(app.Trans("user.create.validation.password"))

	if rv.HasErrors() {
		rv.Keep()
		return rv.ErrorMap()
	}

	return nil
}
