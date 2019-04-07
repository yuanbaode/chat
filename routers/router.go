package routers

import (
	"mychatroom/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	//beego.Router("/index", &controllers.ChatController{}, "*:Index")
	//beego.Router("/login", &controllers.ChatController{}, "post:Login")
	//beego.Router("/chat", &controllers.ChatController{}, "*:Chat")
	//beego.Router("/custom-chat", &controllers.ChatController{}, "get:GetChats")
	//beego.Router("/create-chat", &controllers.ChatController{}, "get:CreateChat")
	beego.Router("/rooms/:user_id", &controllers.RoomController{}, "get:GetRooms")
	beego.Router("/room/:room_id/:user_id", &controllers.RoomController{}, "*:EnterRoom")
	beego.Router("/room/:room_id/:user_id", &controllers.RoomController{}, "*:BroadcastRoom")
}
