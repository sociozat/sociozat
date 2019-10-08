package models

type SettingsModel struct {
	HeaderChannels []UserHeaderChannels `json:"headerChannels"`
	TodaysChannels []UserTodaysChannels `json:"todaysChannels"`
	PostPerPage    int                  `json:"postPerPage"`
	TopicPerPage   int                  `json:"topicPerPage"`
	Theme          string               `json:"theme"`
}
