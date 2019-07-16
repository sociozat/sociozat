package controllers

import (
	"github.com/revel/revel"
)

type AppC struct {
	*revel.Controller
}

func (this AppC) Index() revel.Result {
	title := "Naber"
	return this.Render(title)
}

func (this AppC) testFunc(test string) string {
	return test
}
