package main

import (
	_ "mychatroom/routers"
	"github.com/astaxie/beego"
	"mychatroom/models"
)

func main() {
	//初始化数据库
	models.InitModel()
	beego.BConfig.WebConfig.Session.SessionOn = true
	//beego.BConfig.Listen.EnableAdmin=true
	beego.Run()
}
