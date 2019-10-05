package models

type SettingsModel struct {
	HeaderChannels []UserHeaderChannels `json:"headerChannels"`
	TodaysChannels []UserTodaysChannels `json:"todaysChannels"`
	PostPerPage    int                  `json:"postPerPage"`
	Theme          string               `json:"theme"`
}
