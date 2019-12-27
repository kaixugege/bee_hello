package main

import (
	_ "bee_hello/routers"
	_ "github.com/go-sql-driver/mysql" // import your used driver
	_ "bee_hello/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"bee_hello/models"
	"fmt"
	"strings"
)

func init() {
	//orm.RegisterDriver("mysql", orm.DRMySQL)
	//set default database
	//orm.RegisterDataBase("default", "mysql", "root:@tcp(127.0.0.1:3306)/bee?charset=utf8", 30)
	//orm.RunSyncdb("default", false, true)
}
func main() {
	o := orm.NewOrm()
	o.Using("default") // 默认使用 default，你可以指定为其他数据库

	profile := new(models.Profile)
	profile.Age = 30

	user := new(models.User)
	user.Profile = profile
	user.Name = "slene"

	fmt.Println(o.Insert(profile))
	fmt.Println(o.Insert(user))

	initTemplate()
	beego.Run()
}

func initTemplate() {
	beego.AddFuncMap("equrl", func(x, y string) bool {
		s1 := strings.Trim(x, "/")
		s2 := strings.Trim(y, "/")
		fmt.Println(s1,"  ",s2)
		return strings.Compare(s1, s2) == 0
	})
}