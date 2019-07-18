package models

import (
	"github.com/twinj/uuid"
)

type PostModel struct {
	ID          uuid.UUID
	Title       string
	Description string
	User        UserModel
}
