package controllers

import (
	"github.com/astaxie/beego"
	"fmt"
)


// 约定：如果子controller 存在NestPrepare()方法，就实现了该接口，
type NestPreparer interface {
	NestPreparer()
}


type BaseController struct {
	beego.Controller
}
//  Prepare 方法是每次请求都要调用的
func (ctx *BaseController) Prepare() {
	fmt.Println("Prepare  当前url为",ctx.Ctx.Request.RequestURI)
	// 将页面路径 保存到 Path变量里面
	ctx.Data["Path"] = ctx.Ctx.Request.RequestURI
	fmt.Println("Prepare  当前url为 - ",	ctx.Data["Path"] )

	// 判断子类是否实现了NestPreparer接口，如果实现了就调用接口方法。
	if app,ok := ctx.AppController.(NestPreparer);ok{
		app.NestPreparer()
	}
}
