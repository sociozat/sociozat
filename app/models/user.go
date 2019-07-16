package models

import (
	"github.com/twinj/uuid"
	"github.com/revel/revel"
)

type User struct {
	ID uuid.UUID
	Name string
	Username string
	Email string
}

func (u User) String() string {
	return fmt.Sprintf("User(%s)", u.Name)
}

