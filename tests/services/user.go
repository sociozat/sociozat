package tests

import (
	"github.com/revel/revel/testing"
	"sozluk/app/services"
)

type UserServiceTest struct {
	testing.TestSuite
	UserService services.UserService
}

func (t *UserServiceTest) TestCreateUserValidation() {

}
