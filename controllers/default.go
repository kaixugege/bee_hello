package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.html"
}

type EditorController struct {
	beego.Controller
}

func (e *EditorController) Get() {
	e.TplName = "editor.html"
}

type TestContronller struct {
	beego.Controller
}

func (c *TestContronller) SelfTest() {
	c.Ctx.WriteString("this is myself  controller!")
}

func (c *TestContronller) GetData() {
	id := c.GetString("id")
	c.Ctx.WriteString(id)
	name := c.Input().Get("name")
	c.Ctx.WriteString(name)
}