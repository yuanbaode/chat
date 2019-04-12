package controllers

import (
	"mychatroom/services/room"
	"strconv"
	"mychatroom/log"
	"mychatroom/errs"
	"github.com/gorilla/websocket"
	"net/http"
)

var Upgrade websocket.Upgrader

func init() {
	Upgrade = websocket.Upgrader{
		//允许跨域
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
}

type RoomController struct {
	MainController
}

func (c *RoomController) GetRooms() {
	rs := room.NewRoomService()
	rooms := rs.GetRooms()
	c.Success(rooms)

}

func (c *RoomController) BroadcastRoom(){
	var err error
	var data interface{}
	defer func() {
		if err != nil {
			c.Error(err)
		} else {
			c.Success(data)
		}

	}()
	rId := c.Ctx.Input.Query(":room_id")
	roomId, err := strconv.Atoi(rId)
	if err != nil {
		log.Error("invalid input data, roomId :%s illegal . ", rId)
		return
	}
	uId := c.Ctx.Input.Query(":user_id")
	userId, err := strconv.Atoi(uId)
	if err != nil {
		log.Error("invalid input data, userId :%s illegal . ", rId)
		return
	}
	if userId!=1{
		log.Error("not admin .")
		return
	}
	conn, err := Upgrade.Upgrade(c.Ctx.ResponseWriter, c.Ctx.Request, nil)
	if err != nil {
		log.Error("new ws connect faild .  err:%s\n", err.Error())
		err = errs.WS_CONNECT_FAILD
		return
	}
	defer func() {
		conn.Close()
	}()
	rs := room.NewRoomService()
	err = rs.BroadcastRoom(int64(roomId), int64(userId), conn)
}
func (c *RoomController) EnterRoom() {
	var err error
	var data interface{}
	defer func() {
		if err != nil {
			c.Error(err)
		} else {
			c.Success(data)
		}

	}()
	rId := c.Ctx.Input.Query(":room_id")
	roomId, err := strconv.Atoi(rId)
	if err != nil {
		log.Error("invalid input data, roomId :%s illegal . ", rId)
		return
	}
	uId := c.Ctx.Input.Query(":user_id")
	userId, err := strconv.Atoi(uId)
	if err != nil {
		log.Error("invalid input data, userId :%s illegal . ", rId)
		return
	}
	conn, err := Upgrade.Upgrade(c.Ctx.ResponseWriter, c.Ctx.Request, nil)
	if err != nil {
		log.Error("new ws connect faild .  err:%s\n", err.Error())
		err = errs.WS_CONNECT_FAILD
		return
	}
	defer func() {
		conn.Close()
	}()
	rs := room.NewRoomService()
	err = rs.EnterRoom(int64(roomId), int64(userId), conn)

}

