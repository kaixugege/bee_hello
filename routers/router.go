package routers

import (
	"bee_hello/controllers"
	"github.com/astaxie/beego"
)

func init() {
	//beego.Router("/", &controllers.IndexController{})
	//beego.Router("/editor", &controllers.EditorController{})
	beego.Include(
		&controllers.UserController{},
		&controllers.IndexController{})
	beego.Router("/getdata", &controllers.TestContronller{}, "get:GetData")
	beego.ErrorController(&controllers.ErrorController{})
}
