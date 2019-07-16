package controllers

import (
	"github.com/revel/revel"
	"sozluk/app/models"
)

type UserC struct {
	AuthC
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

	//save to db

	//set success flash
	this.Flash.Success("Welcome," + user.Name)
	return this.Redirect(AppC.Index)
}
