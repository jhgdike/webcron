package main

import (
	"github.com/astaxie/beego"
	"github.com/lisijie/webcron/app/controllers"
	"github.com/lisijie/webcron/app/jobs"
	_ "github.com/lisijie/webcron/app/mail"
	"github.com/lisijie/webcron/app/models"
	"html/template"
	"net/http"
)

const VERSION = "1.0.0"

func main() {
	models.Init()
	jobs.InitJobs()

	// 设置默认404页面
	beego.ErrorHandler("404", func(rw http.ResponseWriter, r *http.Request) {
		t, _ := template.New("404.html").ParseFiles(beego.BConfig.WebConfig.ViewsPath + "/error/404.html")
		data := make(map[string]interface{})
		data["content"] = "page not found"
		t.Execute(rw, data)
	})

	// 生产环境不输出debug日志
	if beego.AppConfig.String("runmode") == "prod" {
		beego.SetLevel(beego.LevelInformational)
	}
	beego.AppConfig.Set("version", VERSION)

	// 路由设置
	beego.Router("/cron", &controllers.MainController{}, "*:Index")
	beego.Router("/cron/login", &controllers.MainController{}, "*:Login")
	beego.Router("/cron/logout", &controllers.MainController{}, "*:Logout")
	beego.Router("/cron/profile", &controllers.MainController{}, "*:Profile")
	beego.Router("/cron/gettime", &controllers.MainController{}, "*:GetTime")
	beego.Router("/cron/help", &controllers.HelpController{}, "*:Index")
	//beego.AutoRouter(&controllers.TaskController{})
	//beego.AutoRouter(&controllers.GroupController{})
	beego.AutoPrefix("/cron/", &controllers.TaskController{})
	beego.AutoPrefix("/cron/", &controllers.GroupController{})

	beego.BConfig.WebConfig.Session.SessionOn = true
	beego.Run()
}
