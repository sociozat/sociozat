package controllers

import (
	"sociozat/app/services"
	"sociozat/app/websocket"
	"strconv"
	"strings"

	"github.com/revel/revel"
)

//Post struct
type Post struct {
	App
	PostService services.PostService
}

//New renders new post template
func (c Post) New() revel.Result {
	u := c.connected()
	if u == nil {
		c.Flash.Error(c.Message("auth.login.required"))
		return c.Redirect(App.Index)
	}

	title := c.Message("post.new.title")
	return c.Render(title)
}

//Create handles POST request to create new post with topic
func (c Post) Create(name string, content string, tags string) revel.Result {
	u := c.connected()
	if u == nil {
		c.Flash.Error(c.Message("auth.login.required"))
		return c.Redirect(Post.New)
	}

	c.PostService.Validation = c.Validation //this should come from controller, cuz of pointed
	post, v, err := c.PostService.CreateNewPost(name, content, tags, u)

	if err != nil {
		c.FlashParams()
		return c.Redirect(Post.Create)
	}

	if v != nil {
		c.FlashParams()
		return c.Redirect(Post.New)
	}

    //publish via websocket
    var list []string
    for _, c := range post.Topic.Channels {
        list = append(list, strconv.FormatUint(uint64(c.ID), 10))
    }
    websocket.Publish("channels", strings.Join(list, ","))


	c.Flash.Success(c.Message("topic.create.success.message"))
	return c.Redirect(Post.View, post.ID)
}

//View renders post by id
func (c Post) View(id int) revel.Result {
	post, err := c.PostService.FindByID(id)

	if err != nil {
		c.FlashParams()
		return c.Redirect(App.Index)
	}

	var title = c.Message("post.single.title", post.User.Username, post.Topic.Name)
	var topic = post.Topic
	return c.Render(title, post, topic)
}

//Edit post by id
func (c Post) Edit(id int) revel.Result {
	u := c.connected()

	if u == nil {
		c.Flash.Error(c.Message("auth.login.required"))
		return c.Redirect(Post.New)
	}

	post, err := c.PostService.FindByID(id)

	if err != nil {
		c.FlashParams()
		return c.Redirect(App.Index)
	}

	if u.UserID != post.User.UserID {
		c.Flash.Error(c.Message("post.update.denied"))
		return c.Redirect(Post.View, id)
	}

	var title = c.Message("post.single.title", post.User.Username, post.Topic.Name)
	var topic = post.Topic
	return c.Render(title, post, topic)
}

//UPdate post by id
func (c Post) Update(id int) revel.Result {
	u := c.connected()

	if u == nil {
		c.Flash.Error(c.Message("auth.login.required"))
		return c.Redirect(Post.New)
	}

	post, err := c.PostService.FindByID(id)

	if err != nil {
		c.FlashParams()
		return c.Redirect(App.Index)
	}

	if u.UserID != post.User.UserID {
		c.Flash.Error("post.update.denied")
		return c.Redirect(Post.View, id)
	}

	c.PostService.Validation = c.Validation //this should come from controller, cuz of pointed
	post.Content = c.Params.Form.Get("content")
	post, v, err := c.PostService.UpdatePost(post)

	if v != nil {
		c.FlashParams()
		return c.Redirect(Post.Edit, post.ID)
	}

	c.Flash.Success(c.Message("post.update.success.message"))
	return c.Redirect(Post.View, post.ID)
}
