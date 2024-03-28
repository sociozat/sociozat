package models

type SettingsModel struct {
	HeaderChannels []UserHeaderChannels `json:"headerChannels"`
	TrendingChannels []UserTrendingChannels `json:"trendingChannels"`
	PostPerPage    int                  `json:"postPerPage"`
	TopicPerPage   int                  `json:"topicPerPage"`
	Theme          string               `json:"theme"`
}
