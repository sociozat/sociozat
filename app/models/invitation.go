package models

import (
	"github.com/jinzhu/gorm"
)

//
type InvitationModel struct {
	gorm.Model
	UserID      string `gorm:"type:char(36);index:invitation_user"`
	UsedByID    string `gorm:"type:char(36);index:invitation_used_by"`
	Code        string `gorm:"type:char(36);index:invitation_code"`
}

//TableName sets table name on db
func (p InvitationModel) TableName() string {
	return "invitations"
}
