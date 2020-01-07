package controllers

import (
	"bee_hello/models"
	"bee_hello/syserrors"
	"errors"
	"fmt"
	"github.com/astaxie/beego/logs"
	"strings"
)

type UserController struct {
	BaseController
}

// @router /login [post]
func (c *UserController) Login() {

	email := c.GetMustString("email", "邮箱不能为空！")
	pwd := c.GetMustString("password", "密码不能为空！")
	logs.Debug("email", email, "pwd", pwd)
	var (
		user models.User
		err  error
	)
	if user, err = c.Dao.QueryUserByEmailAndPassword(email, pwd); err != nil {
		c.Abort500(syserrors.NewError("邮箱或密码不对", err))
	}
	c.SetSession(SESSION_USER_KEY, user)
	c.JSONOk("登陆成功", "/")
	logs.Debug("登陆成功", email, pwd)
}

func (ctx *BaseController) Abort500(err error) {
	ctx.Data["err"] = err
	ctx.Abort("500")
}

//	@router /logout [get]
func (c *UserController) Logout() {

	c.MustLogin()                  // 必须登陆
	c.DelSession(SESSION_USER_KEY) //删除session
	c.Redirect("/", 302)           //跳转路由

}

// @router /setting/editor [post]
func (c *UserController) Editor() {
	editor := c.GetMustString("editor", "default")
	if !strings.EqualFold(editor, "markdown") {
		editor = "default"
	}
}

func init() {

}

// @router /reg [post]
func (c *UserController) Reg() {
	name := c.GetMustString("name", "昵称不能为空")
	email := c.GetMustString("email", "邮箱不能为空！")
	pwd1 := c.GetMustString("password", "密码不能为空！")
	pwd2 := c.GetMustString("password2", "确认密码不能为空！")
	fmt.Printf("POST Reg : name %s, email %s, pwd1 %s, pwd2", name, email, pwd1, pwd2)
	logs.Alert("Reg : name %s, email %s, pwd1 %s, pwd2 %s", name, email, pwd1, pwd2)

	if strings.Compare(pwd1, pwd2) != 0 {
		c.Abort500(errors.New("密码与确认密码 必须要一致"))
	}

}
