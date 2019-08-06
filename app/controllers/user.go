package controllers

import (
	"sozluk/app"
	"sozluk/app/models"
	"sozluk/app/services"

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
		c.Flash.Error(err.Error())
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
			c.Session["user"] = string(username)
			c.Session["fulluser"] = user

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

func (c User) Settings() revel.Result {

	return c.Render()
}

//Logout user
func (c User) Logout() revel.Result {
	for k := range c.Session {
		delete(c.Session, k)
	}
	return c.Redirect(App.Index)
}
