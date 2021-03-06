package models

import (
	"time"
	"github.com/astaxie/beego/orm"
	"strings"
	"regexp"
	"strconv"
	"fmt"
	"errors"
)

type MessageType int

const (
	LOGIN     MessageType = 1
	LOGOUT    MessageType = 2
	WHISPER   MessageType = 3
	GROUPCHAT MessageType = 4
	SYSTEM MessageType = 5
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
	Result     string    `json:"text" orm:"size(15);column(text)"`
	Amount     int64     `json:"amount"`
	UserId     *User     `json:"user_id" orm:"rel(fk);column(user_id)"`
	CreateTime time.Time `json:"create_time" orm:"auto_now_add;type(datetime)"`
}

func (i *InfoStored) TableName() string {
	return "lottery_log"
}
func (i *InfoStored) Insert(xOrm orm.Ormer) (err error) {
	_, err = xOrm.Insert(i)
	return
}
//func (i *InfoStored) Stored(xOrm orm.Ormer)(err error){
//
//}
func UnmarshalInfo(info []byte, infoStored *InfoStored) error {
	s := string(info)
	fmt.Println(s)
	splitsInfo := strings.Split(s, "\n")
	re := regexp.MustCompile(`[1-9]{1}[0-9]*`)
	periodS := re.FindAllString(splitsInfo[0], -1)
	period, err := strconv.Atoi(periodS[0])
	if err != nil {
		return err
	}
	infoStored.Period = int64(period)
	if len(splitsInfo) < 2 {
		return errors.New("splitsInfo length < 2. ")
	}
	splitsInfo2 := strings.Split(splitsInfo[1], "/")
	infoStored.Result = splitsInfo2[0]
	if len(splitsInfo2) < 2 {
		return errors.New("splitsInfo2 length < 2. ")
	}
	amount, err := strconv.Atoi(splitsInfo2[1])
	if err != nil {
		return err
	}
	infoStored.Amount = int64(amount)
	return nil
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
