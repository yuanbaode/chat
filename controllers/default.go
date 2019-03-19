package controllers

import (
	"github.com/astaxie/beego"
	"mychatroom/models"
	"mychatroom/errs"
	"reflect"
	"net/http"
	"strconv"
	"fmt"
)

type MainController struct {
	beego.Controller
	User *models.User
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
