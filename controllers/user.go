package controllers

import (
	"bee_hello/models"
	"bee_hello/syserrors"
	"strings"
)

type UserController struct {
	BaseController
}

// @router /login [post]
func (c *UserController) Login() {
	/*c.GetMustString(agr0,arg1 string ) 是在BaseController里面定义的，第一个参数获取请求的参数的键值对的key，请求后，
	如key对于的value是空，就返回第二个参数*/
	email := c.GetMustString("email", "邮箱不能为空！")
	pwd := c.GetMustString("password", "密码不能为空！")
	var (
		user *models.User
		err  error
	)
	if user, err = models.QueryUserByEmailAndPassword(email, pwd); err != nil {
		c.Abort500(syserrors.NewError("邮箱或密码不对", err))
	}
	c.SetSession(SESSION_USER_KEY, user)
	c.JSONOk("登陆成功", "/")
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
