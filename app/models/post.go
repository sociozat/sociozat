package models

import (
	"github.com/twinj/uuid"
)

type PostM struct {
	ID          uuid.UUID
	Title       string
	Description string
	User        UserModel
}
