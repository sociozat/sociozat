package models

import (
	"time"

	"github.com/gosimple/slug"
	"github.com/jinzhu/gorm"
)

//ChannelModel struct
type ChannelModel struct {
	gorm.Model
	Name      string
	Slug      string     `gorm:"index:channel_slug"`
	CreatedAt *time.Time `gorm:"-" json:"-"`
	UpdatedAt *time.Time `gorm:"-" json:"-"`
	DeletedAt *time.Time `gorm:"-" json:"-"`
}

//TableName sets table name on db
func (c ChannelModel) TableName() string {
	return "channels"
}

//NewChannel returns a new ChannelModelInstance
func NewChannel(name string) ChannelModel {
	return ChannelModel{Name: name, Slug: slug.Make(name)}
}
