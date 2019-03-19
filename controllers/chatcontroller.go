package controllers

import (
	"mychatroom/services/user"
	"mychatroom/log"
	"mychatroom/models"
	"fmt"
)

type ChatController struct {
	MainController
}

func (c *ChatController) Chat() {

	c.TplName = "chat.html"
}
func (c *ChatController) Index() {
	callBack := c.Ctx.Input.Query("callback")
	c.Data["callback"] = callBack
	c.TplName = "login.html"
}
func (c *ChatController) Login() {
	callBack := c.Ctx.Input.Query("callback")
	username := c.Ctx.Input.Query("username")
	password := c.Ctx.Input.Query("password")
	fmt.Println(callBack, username, password)
	var err error
	s := user.NewUserService()
	if err = s.Login(username, password); err != nil {
		log.Error("Login err:%s\n", err.Error())
		c.Redirect("/index", 301)
	} else {
		se := c.Ctx.GetCookie("beegosessionID")
		c.SetSession(se, s.Auth)
	}
	if callBack != "" {
		log.Info("callBack is %s\n", callBack)
		c.Redirect(callBack, 200)
	} else {

		c.TplName = "admin.html"
	}
}
func (c *ChatController) Prepare() {
	if c.Ctx.Input.URL() != "/login" && c.Ctx.Input.URL() != "/index" {
		s := c.Ctx.GetCookie("beegosessionID")
		log.Info("beegosessionID is %s\n", s)
		ifLogin := c.GetSession(s)
		if ifLogin == nil {
			log.Error("ifLogin = nil .")
			c.Redirect("/index?callback="+c.Ctx.Input.URL(), 301)
		}
		var user *models.User
		var ok bool

		if user, ok = ifLogin.(*models.User); !ok {
			log.Error("ifLogin not User .")
			c.Redirect("/index?callback="+c.Ctx.Input.URL(), 301)
		}
		c.User = user

	}

}


func (c *ChatController)CreateChat(){

}