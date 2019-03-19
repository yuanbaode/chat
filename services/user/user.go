package user

import (
	"mychatroom/models"
	"github.com/astaxie/beego/orm"
	"mychatroom/log"
	"crypto/sha256"
	"io"
	"mychatroom/errs"
	"encoding/hex"
)

type UserService struct {
	Auth *models.User
}

func NewUserService() *UserService {
	return new(UserService)
}

func (s *UserService) Login(userName, passwd string) (err error) {
	user := new(models.User)
	xOrm := orm.NewOrm()
	if _, err = user.GetUserByUserId(xOrm, userName); err != nil {
		log.Error("GetUserByUserId err:%s\n", err.Error())
		return
	}
	sha := sha256.New()
	io.WriteString(sha, passwd)
	shaPasswd := sha.Sum(nil)
	if user.Password ==  hex.EncodeToString(shaPasswd) {
		s.Auth = user
		return nil
	}
	return errs.Permission_Deny
}
