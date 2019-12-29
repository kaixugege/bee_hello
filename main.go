package main

import (
	_ "bee_hello/routers"
	_ "github.com/go-sql-driver/mysql" // import your used driver
	_ "bee_hello/models"
	"github.com/astaxie/beego"
	"fmt"
	"strings"
	"encoding/gob"
	"bee_hello/models"
)

func init() {
	//orm.RegisterDriver("mysql", orm.DRMySQL)
	//set default database
	//orm.RegisterDataBase("default", "mysql", "root:@tcp(127.0.0.1:3306)/bee?charset=utf8", 30)
	//orm.RunSyncdb("default", false, true)
}
func main() {
	initSession()
	//o := orm.NewOrm()
	//o.Using("default") // 默认使用 default，你可以指定为其他数据库
	//
	//profile := new(models.Profile)
	//profile.Age = 30
	//
	//user := new(models.User)
	//user.Profile = profile
	//user.Name = "slene"
	//
	//fmt.Println(o.Insert(profile))
	//fmt.Println(o.Insert(user))

	initTemplate()
	beego.Run()
}

func initSession() {
	//beego的session序列号是用gob的方式，因此需要将注册models.User
	gob.Register(models.User{})
	//https://beego.me/docs/mvc/controller/session.md
	beego.BConfig.WebConfig.Session.SessionOn = true
	beego.BConfig.WebConfig.Session.SessionName = "liteblog-key"
	beego.BConfig.WebConfig.Session.SessionProvider = "file"
	beego.BConfig.WebConfig.Session.SessionProviderConfig = "data/session"
}

func initTemplate() {
	beego.AddFuncMap("equrl", func(x, y string) bool {
		s1 := strings.Trim(x, "/")
		s2 := strings.Trim(y, "/")
		fmt.Println(s1, "  ", s2)
		return strings.Compare(s1, s2) == 0
	})
}
