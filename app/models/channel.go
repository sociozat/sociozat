package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type ChannelModel struct {
	gorm.Model
	Name      string
	Slug      string     `gorm:"index:channel_slug"`
	CreatedAt *time.Time `gorm:"-"`
	UpdatedAt *time.Time `gorm:"-"`
	DeletedAt *time.Time `gorm:"-"`
}

func (c ChannelModel) TableName() string {
	return "channels"
}
