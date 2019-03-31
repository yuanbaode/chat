package models

import (
	"github.com/astaxie/beego/orm"
	"mychatroom/log"
	_ "github.com/go-sql-driver/mysql"
	"os"
)
var ORM orm.Ormer
func InitModel() {
	var err error
	//sqlconn := "root:Root#123@tcp(118.25.102.107:3306)/dev_chatroom?charset=utf8"
	sqlconn := "root:Root#123@tcp(118.25.102.107:3306)/player_api?charset=utf8&loc=Local"
	if err = orm.RegisterDriver("mysql", orm.DRMySQL); err != nil {
		log.Error("database register driver err:%s\n", err.Error())
		os.Exit(1)
	}
	err = orm.RegisterDataBase("default", "mysql", sqlconn)
	if err != nil {
		log.Error("database connect err:%s\n", err.Error())
		os.Exit(1)
	}
	//orm.DefaultTimeLoc = time.Local
	//orm.RegisterModel(new(Room))
	orm.RegisterModel(new(User))
	orm.RegisterModel(new(InfoStored))

	if err=orm.RunSyncdb("default", false, true);err!=nil{
		log.Error("database connect err:%s\n", err.Error())
		os.Exit(1)
	}
	ORM=orm.NewOrm()

}
