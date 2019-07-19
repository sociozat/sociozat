package controllers

import (
	"encoding/json"
	"github.com/revel/revel"
	"sozluk/app/models"
	"sozluk/app/services"
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

	u := models.UserModel{}

	if c.Session["user"] == nil {
		return nil
	}
	if err := json.Unmarshal([]byte(c.Session["user"].(string)), &u); err == nil {
		return &u
	}

	return nil
}

//GetUser gets user by username
func (c App) GetUser(username string) *models.UserModel {

	slug := revel.Slug(username)
	user, err := c.UserService.GetBySlug(slug)

	if err == nil {
		return user
	}

	return nil
}
