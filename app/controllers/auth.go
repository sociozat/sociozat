package controllers

import (
	M "sozluk/app/models"
)

type Auth struct {
	App
}

func (this Auth) handle() {
	// do authentication here
}

func (this Auth) CurrentUser() M.UserModel {
	//sample user instance
	m := M.NewUser("John", "HiJohn", "john@doe.com", "12345")
	// m := models.UserM{uuid.NewV4(), "John", "HiJohn", "john@doe.com"}
	return m
	// u := User{uuid.UUID, "John", "HiJohn", "john@doe.com"}
}
