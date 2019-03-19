package models

import (
	"time"
	"github.com/astaxie/beego/orm"
)

type MessageType int

const (
	LOGIN     MessageType = 1
	LOGOUT    MessageType = 2
	WHISPER   MessageType = 3
	GROUPCHAT MessageType = 4
)

type Message struct {
	MType MessageType
	From  User
	Msg   interface{}
}
type Info struct {
	Period int64  `json:"period"`
	Result string `json:"result"`
	Amount int64  `json:"amount"`
}

type InfoStored struct {
	Id         int64     `json:"id" orm:"pk;auto"`
	RoomId     int64     `json:"room_id" `
	Period     int64     `json:"period"`
	Result     string    `json:"result" orm:"size(15)"`
	Amount     int64     `json:"amount"`
	UserId     *User     `json:"user_id" orm:"rel(fk);column(user_id)"`
	CreateTime time.Time `json:"create_time" orm:"auto_now_add;type(datetime)"`
}

func (i *InfoStored) TableName() string {
	return "info_stored"
}
func (i *InfoStored) Insert(xOrm orm.Ormer) (err error) {
	_, err = xOrm.Insert(i)
	return
}

type HistoryMessage struct {
	Id         int64     `json:"id" orm:"pk"`
	RoomId     int64     `json:"room_id" `
	Message    string    `json:"message" orm:"text"`
	UserId     *User     `json:"user_id" orm:"rel(fk)"`
	CreateTime time.Time `json:"create_time" orm:"auto_now_add;type(datetime)"`
}

func (i *HistoryMessage) TableName() string {
	return "history_message"
}
func (i *HistoryMessage) Insert(xOrm orm.Ormer) (err error) {
	_, err = xOrm.Insert(i)
	return
}
