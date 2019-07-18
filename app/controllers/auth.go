package controllers

import (
	"sozluk/app/models"

	"github.com/revel/revel"
)

// Auth struct
type Auth struct {
	App
}

func (c Auth) handle() {
	// do authentication here
}

func (c Auth) connected() *models.UserModel {
	if c.ViewArgs["user"] != nil {
		return c.ViewArgs["user"].(*models.UserModel)
	}
	if username, ok := c.Session["user"]; ok {
		return c.GetUser(username.(string))
	}

	return nil
}

//GetUser gets user by username
func (c Auth) GetUser(username string) *models.UserModel {

	slug := revel.Slug(username)
	user, err := c.UserService.GetBySlug(slug)

	if err == nil {
		return user
	}

	return nil
}
