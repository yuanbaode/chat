package log

import (
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego"
)

var logger *logs.BeeLogger

func init() {
	logger = logs.GetBeeLogger()
	logger.SetLevel(beego.LevelDebug)
	logger.EnableFuncCallDepth(true)
	logger.SetLogger(logs.AdapterFile,`{"filename":"chat.log","level":7,"maxlines":0,"maxsize":0,"daily":true,"maxdays":10,"color":true}`)
}

func Warn(format string, v ...interface{}) {
	logger.Warn(format, v)
}
func Error(format string, v ...interface{}) {
	logger.Error(format, v)
}
func Info(format string, v ...interface{}) {
	logger.Info(format, v)
}
func Debug(format string, v ...interface{}) {
	logger.Debug(format, v)
}
