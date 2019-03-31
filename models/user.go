package models

import (
	"github.com/astaxie/beego/orm"
	"mychatroom/log"
	"mychatroom/errs"
)

type User struct {
	Id     int64  `orm:"pk" json:"id"`
	Name   string `orm:"size(20);column(username)" json:"name"`
	UserId string `orm:"-" json:"user_id"`
	//CurrentRoom *Room  `orm:"rel(fk);column(current_room_id)" json:"current_room"`
	CurrentRoom int64  `orm:"-" json:"current_room"`
	Password    string ` json:"password"`
	AvatarUrl   string `json:"avatar_url" orm:"column(avatar_url)"`
}

func TableName() string {
	return "user"
}

func (user *User) GetUserById(xOrm orm.Ormer, id int64) (*User, error) {
	user.Id = id
	var err error
	if err = xOrm.Read(user, "id"); err != nil {
		log.Error("database operate err:%s\n", err.Error())
		return nil, err
	}
	return user, nil
}
func (user *User) GetUserByUserId(xOrm orm.Ormer, userId string) (*User, error) {
	user.UserId = userId
	var err error
	if err = xOrm.Read(user, "user_id"); err != nil {
		log.Error("database operate err:%s\n", err.Error())
		return nil, errs.DB_OPERATR_ERROR
	}
	return user, nil
}

func (user *User) UpdateUserById(xOrm orm.Ormer, userId int64, column ... string) (err error) {
	user.Id = userId
	if _, err = xOrm.Update(user, column ...); err != nil {
		log.Error("database operate err:%s\n", err.Error())
	}
	return
}
