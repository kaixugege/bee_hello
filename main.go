package main

import (
	"bee_hello/models"
	_ "bee_hello/models"
	_ "bee_hello/routers"
	"encoding/gob"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	_ "github.com/go-sql-driver/mysql" // import your used driver
	"os"
	"strings"
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
	initLog()
	initTemplate()
	beego.Run()
}
func initLog() {
	if err := os.MkdirAll("data/logs", 0777); err != nil {
		beego.Error(err)
		return
	}

	// 设置配置文件
	//jsonConfig := `{
	//    "filename" : "/logs/test.log", // 文件名
	//    "maxlines" : 2000,       // 最大行
	//    "maxsize"  : 10240       // 最大Size
	//}`
	//logs.SetLogger("file", jsonConfig) // 设置日志记录方式：本地文件记录

	logs.SetLogger("file", `{"filename":"data/logs/lyblog.log","level":7,"maxlines":0,"maxsize":0,"daily":true,"maxdays":10}`)
	logs.Async(1e3)
	logs.SetLevel(logs.LevelDebug) // 设置日志写入缓冲区的等级
	logs.EnableFuncCallDepth(true) // 输出log时能显示输出文件名和行号（非必须）

	//logs.SetLogger(logs.AdapterFile,`{"filename":`+beego.AppConfig.String("log_path")+`"/project.log","level":7,"maxlines":0,"maxsize":0,"daily":true,"maxdays":10}`)
	//logs.SetLogger(logs.AdapterFile,`{"filename":"`+beego.AppConfig.String("log_path")+`/project.log","level":7,"maxlines":0,"maxsize":0,"daily":true,"maxdays":10}`)
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
