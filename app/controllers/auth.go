package controllers

import (
	"github.com/twinj/uuid"
	"github.com/bencagri/sozluk/app/models"
)

type Auth struct {
	App
}

func (this Auth) handle() {
	// do authentication here
}

func (this Auth) currentUser() {
	u :=  User{uuid.UUID, "John", "HiJohn", "john@doe.com"}
}