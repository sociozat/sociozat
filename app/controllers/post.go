package controllers

import (
	"github.com/revel/revel"
	"sozluk/app"
	"sozluk/app/services"
)

//Post struct
type Post struct {
	App
	PostService services.PostService
}

func (c Post) New() revel.Result {
	if user := c.connected(); user == nil {
		c.Flash.Error(app.Trans("auth.login.required"))
		return c.Redirect(App.Index)
	}

	title := app.Trans("post.new.title")
	return c.Render(title)
}

//Create handles POST request to create new post with topic
func (c Post) Create(name string, content string) revel.Result {
	u := c.connected()
	if u == nil {
		c.Flash.Error(app.Trans("auth.login.required"))
		return c.Redirect(Post.New)
	}
	c.PostService.Validation = c.Validation //this should come from controller, cuz of pointed
	post, v, err := c.PostService.CreateNewPost(name, content, u)

	if err != nil {
		c.FlashParams()
		return c.Redirect(Post.Create)
	}

	if v != nil {
		c.FlashParams()
		return c.Redirect(Post.New)
	}

	return c.Redirect(Post.View, post.ID)
}

func (c Post) View(id uint) revel.Result {
	return c.Render(id)
}
