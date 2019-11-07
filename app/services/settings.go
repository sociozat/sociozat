package services

import (
	"encoding/json"
	"sociozat/app"
	"sociozat/app/models"
	"sociozat/app/repositories"
	"time"

	"github.com/mitchellh/mapstructure"
	"github.com/revel/revel"
)

var perPage = map[string]int{
	"1": 5,
	"2": 10,
	"3": 25,
}

type SettingsService struct {
	ChannelRepository repositories.ChannelRepository
	UserRepository    repositories.UserRepository
	PostService       PostService
}

//MapSettings generates SettingsModel instance from json
func (s SettingsService) MapSettings(sets interface{}) (models.SettingsModel, error) {
	settings := models.SettingsModel{}

	err := mapstructure.Decode(sets, &settings)

	return settings, err
}

//TransformValues generates SettingsModel instance from form params
func (s SettingsService) TransformValues(params *revel.Params) models.SettingsModel {
	//header channels
	headerChannels := []models.UserHeaderChannels{}
	if params.Get("header-channels") != "" {
		channels := s.PostService.GenerateChannels(params.Get("header-channels"))
		for _, channel := range channels {
			headerChannel := models.UserHeaderChannels{
				ID:       channel.ID,
				Name:     channel.Name,
				Slug:     channel.Slug,
				LastRead: time.Now(),
			}

			headerChannels = append(headerChannels, headerChannel)
		}
	}

	//todays channels
	todaysChannels := []models.UserTodaysChannels{}
	if params.Get("todays-posts-channels") != "" {
		channels := s.PostService.GenerateChannels(params.Get("todays-posts-channels"))
		for _, channel := range channels {
			todaysChannel := models.UserTodaysChannels{
				ID:   channel.ID,
				Name: channel.Name,
				Slug: channel.Slug,
			}

			todaysChannels = append(todaysChannels, todaysChannel)
		}
	}

	var theme string
	if params.Form.Get("theme") != "" {
		theme = params.Get("theme")
	} else {
		theme = "default"
	}

	settings := models.SettingsModel{
		HeaderChannels: headerChannels,
		TodaysChannels: todaysChannels,
		PostPerPage:    perPage[params.Form.Get("post-per-page")],
		TopicPerPage:   perPage[params.Form.Get("topic-per-page")],
		Theme:          theme,
	}

	return settings

}

//Save update user settings
func (s SettingsService) Save(user *models.UserModel, settings models.SettingsModel) (models.SettingsModel, error) {

	settingsjson, err := json.Marshal(settings)

	if err != nil {
		return settings, err
	}

	_, err = s.UserRepository.UpdateUserColumn(user, "settings", string(settingsjson))

	return settings, err
}

func (s SettingsService) Validate(settings models.SettingsModel, rv *revel.Validation) map[string]*revel.ValidationError {

	rv.Check(settings.PostPerPage,
		revel.Required{},
		revel.MaxSize{1},
		revel.MinSize{3},
	).Message(app.Trans("user.create.validation.password"))

	if rv.HasErrors() {
		rv.Keep()
		return rv.ErrorMap()
	}

	return nil

}
