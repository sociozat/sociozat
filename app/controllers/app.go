package controllers

import (
	"github.com/revel/revel"
)

type App struct {
	*revel.Controller
}

func (this App) Index() revel.Result {
	title := "Naber"
	return this.Render(title)
}

func (this App) testFunc(test string) string {
	return test
}
