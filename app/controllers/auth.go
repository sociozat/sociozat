package controllers

import (
	M "sozluk/app/models"
)

type AuthC struct {
	AppC
}

func (this AuthC) handle() {
	// do authentication here
}

func (this AuthC) CurrentUser() M.UserM {
	//sample user instance
	m := M.NewUser("John", "HiJohn", "john@doe.com", "12345")
	// m := models.UserM{uuid.NewV4(), "John", "HiJohn", "john@doe.com"}
	return m
	// u := User{uuid.UUID, "John", "HiJohn", "john@doe.com"}
}
