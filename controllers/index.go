package controllers

import "github.com/astaxie/beego/logs"

type IndexController struct {
	BaseController
}

//首页
// @router / [get]
func (c *IndexController) Get() {
	c.TplName = "index.html"
}

//留言
// @router /message [get]
func (c *IndexController) GetMessage() {
	c.TplName = "message.html"
}

//关于
// @router /about [get]
func (c *IndexController) GetAbout() {
	c.TplName = "about.html"
}

// @router /user [get]
func (c *IndexController) GetUser() {
	//fmt.Println(c.IsLogin)
	c.TplName = "user.html"
}

// @router /reg [get]
func (c *IndexController) GetReg() {
	logs.Alert("get :: Reg")
	c.TplName = "reg.html"
}
