package controllers

import (
	"github.com/revel/revel"
	"sozluk/app/models"
	"sozluk/app/repositories"
)

type User struct {
	Auth
	UserRepository repositories.UserRepository
}

func (this User) GetUser(id int) revel.Result {
	u := this.CurrentUser()
	return this.RenderJSON(u)
}

func (this User) Register() revel.Result {
	title := "Register"
	return this.Render(title)
}

func (this User) SaveUser(user models.UserModel) revel.Result {
	//validate post data
	user.Validate(this.Validation)

	if this.Validation.HasErrors() {
		this.Validation.Keep()
		this.FlashParams()
		return this.Redirect(User.Register)
	}

	newUser := models.NewUser(user.Username, user.Username, user.Email, user.Password)

	//save to db
	u, err := this.UserRepository.Create(newUser)

	if err != nil {
		this.Flash.Error(err.Error())
		return this.Redirect(User.Register)
	}

	//set success flash
	this.Flash.Success("Welcome," + u.Name)
	return this.Redirect(App.Index)
}
