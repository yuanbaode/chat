package room

import (
	"mychatroom/models"
	"strconv"
	"mychatroom/log"
	"github.com/gorilla/websocket"
	"encoding/json"
	"github.com/astaxie/beego/orm"
)

var ROOMS *models.RoomSet
var RoomForWHISPER *models.RoomSet

func init() {
	rooms := make(map[int64]*models.Room, 100)
	for i := 1; i <= 100; i++ {
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
	Auth models.User
}

func NewRoomService() *RoomService {
	return &RoomService{}
}
func (s *RoomService) GetRooms() map[int64]*models.Room {
	return ROOMS.Data
}
func (s *RoomService) EnterRoom(roomId, userId int64, conn *websocket.Conn) (err error) {
	s.Auth.Id=userId
	client := models.NewClient(s.Auth, conn)
	ROOMS.Lock.RLock()
	room := ROOMS.Data[roomId]
	ROOMS.Lock.RUnlock()
	room.Clients.SET(strconv.Itoa(int(userId)), client)

	room.In <- &models.Message{
		MType: models.LOGIN,
		From:  s.Auth,
		Msg:   "login",
	}
	defer func() {
		room.In <- &models.Message{
			MType: models.LOGOUT,
			From:  s.Auth,
			Msg:   "logout",
		}
	}()
	//启动读协程

	var (
		msgType int
		data    []byte
	)
	xOrm := orm.NewOrm()
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
			info := new(models.Info)
			room.In <- &models.Message{models.GROUPCHAT, client.User, string(data)}
			if err = json.Unmarshal(data, info); err != nil {
				log.Error("unmarshl info error:%s\n", err.Error())
			} else {
				infoStored := new(models.InfoStored)
				infoStored.UserId = &client.User
				infoStored.Amount = info.Amount
				infoStored.Period = info.Period
				infoStored.Result = info.Result
				infoStored.RoomId = roomId
				if err = infoStored.Insert(xOrm); err != nil {
					log.Error("Insert info error:%s\n", err.Error())
				}
			}

		}
		//if err != nil {
		//	break
		//}

	}

	//xOrm := orm.NewOrm()
	//xOrm.Begin()
	//defer func() {
	//	if err != nil {
	//		xOrm.Rollback()
	//	} else {
	//		xOrm.Commit()
	//	}
	//}()
	//user := models.User{Id: userId}
	//if _, err = user.GetUserById(xOrm, userId); err != nil {
	//	log.Error("GetUserById err: %s\n", err.Error())
	//	return
	//}
	//user.CurrentRoom = &models.Room{Id: roomId}
	//if err = user.UpdateUserById(xOrm, userId, "current_room_id"); err != nil {
	//	log.Error("UpdateUserById err: %s\n", err.Error())
	//	return
	//}
	return
}

func (s *RoomService) CreateChat(userId int64, isUser bool, conn *websocket.Conn) (err error) {
	client := models.NewClient(s.Auth, conn)
	RoomForWHISPER.Lock.Lock()
	room, ok := RoomForWHISPER.Data[userId]
	ROOMS.Lock.Unlock()
	if !ok {
		room = models.NewRoom("房间" + strconv.Itoa(int(userId)))
		room.Id = userId
		RoomForWHISPER.Lock.Lock()
		RoomForWHISPER.Data[room.Id] = room
		RoomForWHISPER.Lock.Unlock()
	}
	if isUser {
		room.Clients.SET(strconv.Itoa(int(userId)), client)
	} else {
		room.Clients.SET("Support", client)
	}

	//启动读协程

	var (
		msgType int
		data    []byte
	)
	xOrm := orm.NewOrm()
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
			var user models.User
			if isUser {
				user = client.User
			} else {
				user = models.User{Name: "Support"}
			}
			msg := &models.Message{models.WHISPER, user, string(data)}
			room.In <- msg
			var msgData []byte
			if msgData, err = json.Marshal(msg); err != nil {
				log.Error("marshal info error:%s\n", err.Error())
			} else {
				historyMsg := new(models.HistoryMessage)
				historyMsg.UserId = &client.User
				historyMsg.RoomId = client.User.Id
				historyMsg.Message = string(msgData)
				if err = historyMsg.Insert(xOrm); err != nil {
					log.Error("insert historyMsg error:%s\n", err.Error())
				}
			}

		}
		if err != nil {
			break
		}

	}

	return
}
