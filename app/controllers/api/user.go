package controllers

import (
	"github.com/revel/revel"
)

type User struct {
	Auth
}

func (this User) GetUser(id int) revel.Result {

	return this.Response(id)
}

