package controllers

import (
	"sociozat/app/services"

	"github.com/revel/revel"
)

type Settings struct {
	App
	SettingsService services.SettingsService
}

func (c Settings) View() revel.Result {
	u := c.connected()

	if u == nil {
		c.Flash.Error(c.Message("auth.login.required "))
		c.Redirect(User.Login)
	}

	settings, _ := c.Session.Get("settings")
	uuid := u.UserID
	return c.Render(settings, uuid)
}

func (c Settings) SettingsPost() revel.Result {
	//prepare settings model
	u := c.connected()
	settings := c.SettingsService.TransformValues(c.Params)

	//save
	_, err := c.SettingsService.Save(u, settings)
	if err != nil {
		c.Log.Error("%v", err)
		c.Flash.Error(c.Message("settings.updated.error"))
	} else {
		//override session values
		c.Session.Set("settings", settings)
		c.Flash.Success(c.Message("settings.updated.sucessfully"))
	}

	return c.Redirect(Settings.View)
}
