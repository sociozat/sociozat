package controllers

import (
	"sociozat/app/models"
	"sociozat/app/services"

	"github.com/gosimple/slug"
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

func (c App) connected() *models.UserModel {
	if c.ViewArgs["user"] != nil {
		return c.ViewArgs["user"].(*models.UserModel)
	}
	if username, ok := c.Session["user"]; ok {
		return c.GetUser(username.(string))
	}

	return nil
}

//GetUser gets user by username
func (c App) GetUser(username string) *models.UserModel {

	slug := slug.Make(username)
	user, err := c.UserService.GetBySlug(slug)

	if err == nil {
		return user
	}

	return nil
}
