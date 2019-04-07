package room

import (
	"mychatroom/models"
	"strconv"
	"mychatroom/log"
	"github.com/gorilla/websocket"
	"github.com/astaxie/beego/orm"
)

var ROOMS *models.RoomSet
var RoomForWHISPER *models.RoomSet

func init() {
	rooms := make(map[int64]*models.Room, 100)
	for i := 1; i <= 4; i++ {
		room := models.NewRoom("房间" + strconv.Itoa(i))
		room.Id = int64(i)
		rooms[room.Id] = room
	}
	ROOMS = &models.RoomSet{
		Data: rooms,
	}

	roomsWhisper := make(map[int64]*models.Room, 100)
	RoomForWHISPER = &models.RoomSet{
		Data: roomsWhisper,
	}
}

type RoomService struct {
	Auth *models.User
}

func NewRoomService() *RoomService {
	return &RoomService{}
}
func (s *RoomService) GetRooms() map[int64]*models.Room {
	return ROOMS.Data
}
func (s *RoomService) EnterRoom(roomId, userId int64, conn *websocket.Conn) (err error) {
	userOp := &models.User{}
	s.Auth, err = userOp.GetUserById(models.ORM, userId)
	if err != nil {
		log.Error("getUser error: %s", err.Error())
		return
	}
	client := models.NewClient(*s.Auth, conn)
	ROOMS.Lock.Lock()
	room := ROOMS.Data[roomId]
	ROOMS.Lock.Unlock()
	room.Clients.SET(strconv.Itoa(int(userId)), client)



	var (
		msgType int
		data    []byte
	)

	for {
		msgType, data, err = conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Error("error: %s", err.Error())
			}
			log.Warn("ReadMessage err:%s\n", err.Error())
			room.Clients.SET(strconv.Itoa(int(userId)), nil)
			break

		}
		if msgType == websocket.TextMessage {
			infoStored := new(models.InfoStored)
			if err = models.UnmarshalInfo(data, infoStored); err != nil {
				log.Error("unmarshl info error:%s\n", err.Error())
				continue
			} else {
				xOrm := orm.NewOrm()
				infoStored.UserId = &client.User
				infoStored.RoomId = roomId
				xOrm.Begin()

				if err = infoStored.Insert(xOrm); err != nil {
					log.Error("Insert info error:%s\n", err.Error())
					xOrm.Rollback()
					continue
				}
				user := &models.User{}
				if user, err = user.GetUserById(xOrm, userId); err != nil {
					log.Error("GetUserById  error:%s", err.Error())
					xOrm.Rollback()
					continue
				}
				user.Balance = user.Balance - infoStored.Amount
				if err = user.UpdateUserById(xOrm, userId, "balance"); err != nil {
					log.Error("update balance err:%s", err.Error())
					xOrm.Rollback()
					continue
				}
				xOrm.Commit()
			}
			room.In <- &models.Message{models.GROUPCHAT, client.User, string(data)}

		}


	}

	return
}
func (s *RoomService) BroadcastRoom(roomId, userId int64, conn *websocket.Conn) (err error) {
	ROOMS.Lock.Lock()
	room := ROOMS.Data[roomId]
	ROOMS.Lock.Unlock()
	var (
		msgType int
		data    []byte
	)

	for {
		msgType, data, err = conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Error("error: %s", err.Error())
			}
			log.Warn("ReadMessage err:%s\n", err.Error())
			break
		}
		if msgType == websocket.TextMessage {
			room.In <- &models.Message{models.SYSTEM, models.User{}, string(data)}
		}
	}
	return
}


