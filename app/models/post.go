package models

import (
	"github.com/gosimple/slug"
	"github.com/jinzhu/gorm"
)

type TopicModel struct {
	gorm.Model
	Name      string
	Slug      string         `gorm:"index:topic_slug"`
	Channels  []ChannelModel `gorm:"many2many:topic_channels;"`
	IsLocked  bool           `gorm:"default:'0'"`
	PostCount int
}

func (t TopicModel) TableName() string {
	return "topics"
}

type PostModel struct {
	gorm.Model
	Content  string `gorm:"type:text;"`
	User     *UserModel
	UserID   uint
	Topic    TopicModel
	TopicID  uint
	Likes    int
	Dislikes int
}

//TableName sets table name on db
func (p PostModel) TableName() string {
	return "posts"
}

//CreateNewPost creates a post instance with relations
func CreateNewPost(name string, content string, user *UserModel) *PostModel {

	t := TopicModel{Name: name, Slug: slug.Make(name), IsLocked: false}

	p := PostModel{
		Content: content,
		User:    user,
		Topic:   t,
	}

	return &p
}

//ReplyPost create as post instance with topic and user
func ReplyPost(content string, topic TopicModel, user UserModel) *PostModel {

	p := PostModel{
		Content: content,
		User:    &user,
		Topic:   topic,
	}

	return &p
}
