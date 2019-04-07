package models

import (
	"mychatroom/log"
	"mychatroom/errs"
)

type Token struct {
	Id          int64  `orm:"pk"`
	AccessToken string `json:"access_token"`
	UserId      int64 `json:"user_id"`
}

func (t *Token) TableName() string {
	return "token"
}
func (t *Token) GetToken(userId int) (token string, err error) {
	if err = ORM.QueryTable("token").Filter("user_id", userId).One(t); err != nil {
		log.Error("GetToken err:%s", err.Error())
		err = errs.DB_OPERATR_ERROR
		return
	}
	return t.AccessToken, err
}
