package log

import (
	"github.com/gin-gonic/gin"
	"github.com/lestrrat-go/file-rotatelogs"
	"github.com/pkg/errors"
	"github.com/rifflock/lfshook"
	log "github.com/sirupsen/logrus"
	"io"
	"os"
	"path"
	"runtime"
	"time"
)

// InitLog 初始化日志格式
func InitLog() {
	//日志写入文件时，禁用控制台颜色

	gin.DisableConsoleColor()
	f, _ := os.OpenFile("./new_project", os.O_CREATE|os.O_RDWR|os.O_APPEND, os.ModeAppend) //0644?
	log.SetOutput(f)
	//初始化日志格式
	initLogDetails()
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

}

var Logs = log.New()

func initLogDetails() {
	var logPath = "./"
	baseLogPath := path.Join(logPath, "new_project")

	writer, err := rotatelogs.New(
		baseLogPath+"_%Y%m%d%H%M.log",
		rotatelogs.WithLinkName(baseLogPath), // 生成软链，指向最新日志文件

		rotatelogs.WithMaxAge(time.Hour*24*10), // 文件最大保存时间
		// rotatelogs.WithRotationCount(365),  // 最多存365个文件

		rotatelogs.WithRotationTime(time.Hour*24), // 日志切割时间间隔

	)
	if err != nil {
		Logs.Errorf("config local file system logger error. %+v", errors.WithStack(err))
	}
	Logs.SetLevel(log.DebugLevel)
	Logs.SetReportCaller(true)
	lfHook := lfshook.NewHook(lfshook.WriterMap{
		log.DebugLevel: writer, // 为不同级别设置不同的输出目的
		log.InfoLevel:  writer,
		log.WarnLevel:  writer,
		log.ErrorLevel: writer,
		log.FatalLevel: writer,
		log.PanicLevel: writer,
	}, &log.JSONFormatter{
		PrettyPrint: true,
		CallerPrettyfier: func(frame *runtime.Frame) (function string, file string) {
			//处理文件名字
			fileName := path.Base(frame.File)
			functionName := path.Base(frame.Function)
			return functionName, fileName
		},
	})
	Logs.AddHook(lfHook)
	Logs.Info("log init success")
}
