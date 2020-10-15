package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

type LoginController struct {
	beego.Controller
}

func (c *LoginController) Get() {
	if c.Input().Get("exit") == "true" {
		c.Ctx.SetCookie("uname", "", -1, "/")
		c.Ctx.SetCookie("pwd", "", -1, "/")
		c.Redirect("/", 301)
		return
	}
	c.TplName = "login.html"
}

func (c *LoginController) Post() {
	uname := c.Input().Get("uname")
	pwd := c.Input().Get("pwd")
	autoLogin := c.Input().Get("autoLogin") == "on"

	if uname == beego.AppConfig.String("adminName") && pwd == beego.AppConfig.String("adminPwd") {
		var maxAge int64
		if autoLogin {
			maxAge = 1<<31 - 1
		}
		c.Ctx.SetCookie("uname", uname, maxAge)
		c.Ctx.SetCookie("pwd", pwd, maxAge)
	}
	c.Redirect("/", 301)
	return
}

func checkAccount(c *context.Context) bool {
	ck, err := c.Request.Cookie("uname")
	if err != nil {
		return false
	}
	uname := ck.Value
	ck, err = c.Request.Cookie("pwd")
	if err != nil {
		return false
	}
	pwd := ck.Value
	return uname == beego.AppConfig.String("adminName") && pwd == beego.AppConfig.String("adminPwd")
}
