package controllers

import (
	"sozluk/app/services"

	"github.com/revel/revel"
)

//App struct
type App struct {
	*revel.Controller
	UserService services.UserService
}

//Index renders home page
func (c App) Index() revel.Result {
	title, _ := revel.Config.String("app.name")
	return c.Render(title)
}
