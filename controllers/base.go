package controllers

import (
	"bee_hello/models"
	"bee_hello/syserrors"
	"errors"
	"github.com/astaxie/beego"
)

const SESSION_USER_KEY = "SESSION_USER_KEY"

// 约定：如果子controller 存在NestPrepare()方法，就实现了该接口，
type NestPreparer interface {
	NestPreparer()
}

type BaseController struct {
	beego.Controller

	IsLogin bool        //标识  是否登录
	User    models.User // 登录的用户
}

//  Prepare 方法是每次请求都要调用的
func (ctx *BaseController) Prepare() {
	// 验证用户是否登陆，判断session中是否存在用户，存在就已经登陆，不存在就没有登陆。
	ctx.IsLogin = false
	tu := ctx.GetSession(SESSION_USER_KEY) //从session中读出这个key, 没有的时候是 nil
	if tu != nil {
		//这里要判断类型，
		if u, ok := tu.(models.User); ok {
			ctx.User = u
			ctx.Data["User"] = u
			ctx.IsLogin = true
		}
	}
	ctx.Data["IsLogin"] = ctx.IsLogin

	// 将页面路径 保存到 Path变量里面
	ctx.Data["Path"] = ctx.Ctx.Request.RequestURI

	// 判断子类是否实现了NestPreparer接口，如果实现了就调用接口方法。
	if app, ok := ctx.AppController.(NestPreparer); ok {
		app.NestPreparer()
	}
}

/**

 */
func (ctx *BaseController) GetMustString(key string, msg string) string {
	email := ctx.GetString(key, "")
	if len(email) == 0 {
		ctx.Abort500(errors.New(msg))
	}
	return email
}

type ResultJsonValue struct {
	Code   int         `json:"code"`
	Msg    string      `json:"msg"`
	Action string      `json:"action,omitempty"`
	Count  int         `json:"count,omitempty"`
	Data   interface{} `json:"data,omitempty"`
}

func (ctx *BaseController) JSONOk(msg string, actions ...string) {
	var action string
	if len(actions) > 0 {
		action = actions[0]
	}
	ctx.Data["json"] = &ResultJsonValue{
		Code:   0,
		Msg:    msg,
		Action: action,
	}
	ctx.ServeJSON()
}

/**
必须登陆
*/
func (ctx *BaseController) MustLogin() {
	if !ctx.IsLogin {
		ctx.Abort500(syserrors.NoUserError{}) //如果没有登陆，则抛出没有用户的 error
	}
}
