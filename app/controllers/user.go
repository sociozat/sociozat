package controllers

import (
	"sozluk/app"
	"sozluk/app/models"
	"sozluk/app/routes"
	"sozluk/app/services"

	"github.com/revel/revel"
	"golang.org/x/crypto/bcrypt"
)

//User struct
type User struct {
	Auth
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
		revel.AppLog.Debug(user.HashedPassword)
		err := bcrypt.CompareHashAndPassword([]byte(user.HashedPassword), []byte(password))
		if err == nil {
			c.Session["user"] = user.UserID
			if remember {
				c.Session.SetDefaultExpiration()
			} else {
				c.Session.SetNoExpiration()
			}
			c.Flash.Success("welcome " + username)
			return c.Redirect(App.Index)
		}
	}
	c.Flash.Out["username"] = username
	c.Flash.Error("Login Failed")
	return c.Redirect(App.Index)
}

//Logout user
func (c User) Logout() revel.Result {
	for k := range c.Session {
		delete(c.Session, k)
	}
	return c.Redirect(routes.App.Index())
}
