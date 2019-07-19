package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

//ChannelModel struct
type ChannelModel struct {
	gorm.Model
	Name      string
	Slug      string     `gorm:"index:channel_slug"`
	CreatedAt *time.Time `gorm:"-"`
	UpdatedAt *time.Time `gorm:"-"`
	DeletedAt *time.Time `gorm:"-"`
}

//TableName sets table name on db
func (c ChannelModel) TableName() string {
	return "channels"
}
