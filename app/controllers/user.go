package controllers

import (
	"github.com/revel/revel"
	"sozluk/app/models"
	"sozluk/app/repositories"
)

type UserC struct {
	AuthC
	UserR repositories.UserR
}

func (this UserC) GetUser(id int) revel.Result {
	u := this.CurrentUser()
	return this.RenderJSON(u)
}

func (this UserC) Register() revel.Result {
	title := "Register"
	return this.Render(title)
}

func (this UserC) SaveUser(user models.UserM) revel.Result {
	//validate post data
	user.Validate(this.Validation)

	if this.Validation.HasErrors() {
		this.Validation.Keep()
		this.FlashParams()
		return this.Redirect(UserC.Register)
	}

	newUser := models.NewUser(user.Username, user.Username, user.Email, user.Password)

	//save to db
	u, err := this.UserR.Create(newUser)

	if err != nil {
		this.Flash.Error(err.Error())
		return this.Redirect(UserC.Register)
	}

	//set success flash
	this.Flash.Success("Welcome," + u.Name)
	return this.Redirect(AppC.Index)
}
