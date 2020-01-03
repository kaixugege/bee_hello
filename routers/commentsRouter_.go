package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["bee_hello/controllers:IndexController"] = append(beego.GlobalControllerRouter["bee_hello/controllers:IndexController"],
		beego.ControllerComments{
			Method:           "Get",
			Router:           `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["bee_hello/controllers:IndexController"] = append(beego.GlobalControllerRouter["bee_hello/controllers:IndexController"],
		beego.ControllerComments{
			Method:           "GetAbout",
			Router:           `/about`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["bee_hello/controllers:IndexController"] = append(beego.GlobalControllerRouter["bee_hello/controllers:IndexController"],
		beego.ControllerComments{
			Method:           "GetMessage",
			Router:           `/message`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["bee_hello/controllers:IndexController"] = append(beego.GlobalControllerRouter["bee_hello/controllers:IndexController"],
		beego.ControllerComments{
			Method:           "GetReg",
			Router:           `/reg`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["bee_hello/controllers:IndexController"] = append(beego.GlobalControllerRouter["bee_hello/controllers:IndexController"],
		beego.ControllerComments{
			Method:           "GetUser",
			Router:           `/user`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["bee_hello/controllers:UserController"] = append(beego.GlobalControllerRouter["bee_hello/controllers:UserController"],
		beego.ControllerComments{
			Method:           "Login",
			Router:           `/login`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["bee_hello/controllers:UserController"] = append(beego.GlobalControllerRouter["bee_hello/controllers:UserController"],
		beego.ControllerComments{
			Method:           "Logout",
			Router:           `/logout`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["bee_hello/controllers:UserController"] = append(beego.GlobalControllerRouter["bee_hello/controllers:UserController"],
		beego.ControllerComments{
			Method:           "Reg",
			Router:           `/reg`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["bee_hello/controllers:UserController"] = append(beego.GlobalControllerRouter["bee_hello/controllers:UserController"],
		beego.ControllerComments{
			Method:           "Editor",
			Router:           `/setting/editor`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

}
