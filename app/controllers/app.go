package controllers

import (
	"github.com/revel/revel"
)

type App struct {
	*revel.Controller
}

func (c App) Index() revel.Result {
    title := "Naber"
	return c.Render(title)
}

func (c App) testFunc(test string) string  {
	return test
}