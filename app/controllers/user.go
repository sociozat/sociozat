package controllers

import (
	"encoding/json"
	"sociozat/app"
	"sociozat/app/models"
	"sociozat/app/services"

	"github.com/gosimple/slug"

	"fmt"
	"strconv"

	"github.com/revel/revel"
	"golang.org/x/crypto/bcrypt"
)

//User struct
type User struct {
	App
	UserService services.UserService
}

//Register renderss register route
func (c User) Register() revel.Result {
	title := "Register"
	return c.Render(title)
}

//SaveUser inserts new user to db
func (c User) SaveUser(user models.UserModel) revel.Result {

	//try to create new user
	u, v, err := c.UserService.Create(user, c.Validation)

	if v != nil {
		c.FlashParams()
		return c.Redirect(User.Register)
	}

	if err != nil {
		c.Flash.Error("user exists")
		return c.Redirect(User.Register)
	}
	//set success flash
	c.Flash.Success("Welcome," + u.Name)
	return c.Redirect(App.Index)
}

//Login renderss login route
func (c User) Login() revel.Result {
	title := app.Trans("user.login.title")
	return c.Render(title)
}

//LoginPost via username and password
func (c User) LoginPost(username string, password string, remember bool) revel.Result {
	user := c.GetUser(username)

	if user != nil {
		err := bcrypt.CompareHashAndPassword([]byte(user.HashedPassword), []byte(password))
		if err == nil {
			settings := models.SettingsModel{}
			json.Unmarshal([]byte(user.Settings), &settings)
			user.Settings = ""

			c.Session["user"] = string(username)
			c.Session["fulluser"] = user
			c.Session["settings"] = settings

			if remember {
				c.Session.SetDefaultExpiration()
			} else {
				c.Session.SetNoExpiration()
			}
			c.Flash.Success(c.Message("greeting.user", username))
			return c.Redirect(App.Index)
		}
	}
	c.Flash.Out["username"] = username
	c.Flash.Error(app.Trans("user.login.error"))
	return c.Redirect(App.Index)
}

//Logout user
func (c User) Logout() revel.Result {
	for k := range c.Session {
		delete(c.Session, k)
	}

	//we need defalt settings
	app.SetDefaultSettings(c.Controller)

	return c.Redirect(App.Index)
}

//UserProfile see user profile with posts
func (c User) Profile(username string) revel.Result {
	page, _ := strconv.Atoi(c.Params.Query.Get("page"))
	limit, _ := strconv.Atoi(c.Params.Query.Get("limit"))
	slug := slug.Make(username)

	params := models.SearchParams{
		Page:  page,
		Limit: limit,
		Slug:  slug,
	}
	user, posts, err := c.UserService.GetUserInfo(params)

	//set pages
	c.Params.Query = c.Request.URL.Query()

	var pagination = make(map[int]string)
	for i := 1; i <= posts.TotalPage; i++ {
		c.Params.Query.Set("page", strconv.Itoa(i))
		pageValue := fmt.Sprintf("/u/%s?%s", c.Params.Route.Get("username"), c.Params.Query.Encode())
		pagination[i] = pageValue
	}

	if err != nil {
		c.Flash.Error(app.Trans("user.not.found"))
		return c.Redirect(App.Index)
	}

    canonical := fmt.Sprintf("%s/u/%s",  revel.Config.StringDefault("app.url", ""), c.Params.Route.Get("username"))
	return c.Render(user, canonical, posts, pagination)
}
