package controllers

import (
	"github.com/astaxie/beego"
	"mychatroom/models"
	"mychatroom/errs"
	"reflect"
	"net/http"
	"strconv"
	"fmt"
	"strings"
	"mychatroom/log"
)

type MainController struct {
	beego.Controller
	User *models.User
}

func (c *MainController) Prepare() {
	var err error
	defer func() {
		if err != nil {
			c.Error(err)
		}
	}()
	author := c.Ctx.Input.Header("Authorization")
	log.Info("header is %s" ,c.Ctx.Request.Header)
	if author == "" {
		err = errs.Permission_Deny
		return
	}
	ss := strings.Split(author, " ")
	if len(ss) < 2 {
		err = errs.Permission_Deny
	}
	accessToken := ss[1]
	token := &models.Token{}
	uId := c.Ctx.Input.Query(":user_id")
	userId, err := strconv.Atoi(uId)
	if err != nil {
		log.Error("invalid input data, userId :%s illegal . ", uId)
		return
	}
	var GetAccessToken string
	if GetAccessToken, err = token.GetToken(userId); err != nil {
		log.Error("GetToken err :%s", err.Error())
		err = errs.Permission_Deny
		return
	}
	if GetAccessToken != accessToken {
		err = errs.Permission_Deny
		return
	}
}

func (c *MainController) Success(i interface{}) {
	var m = make(map[string]interface{})
	m["data"] = i
	m["code_status"] = errs.SUCCESS
	c.Ctx.Output.Status = 200
	c.Data["json"] = m
	c.ServeJSON()
}
func (c *MainController) Error(e error) {
	var m = make(map[string]interface{})
	et := reflect.TypeOf(e)
	if et.Name() != "Err" {
		err := errs.NewComplexError(http.StatusBadRequest, e)
		e = err
		m["code_status"] = errs.FAILD
		c.Ctx.Output.Status = err.StatusCode
	} else {
		var i interface{}
		i = e
		s := i.(errs.Err)
		status, _ := strconv.Atoi(string(s)[:3])
		fmt.Println(status, "xxx")
		c.Ctx.Output.Status = status
		m["code_status"] = e
	}
	m["data"] = e.Error()

	c.Data["json"] = m
	c.ServeJSON()
}
