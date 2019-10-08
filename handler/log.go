package handler

import (
	"fmt"
	"os"
	"time"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
)

var GlobalLog *logrus.Entry

func UseLog() *logrus.Entry {
	if GlobalLog == nil {
		fmt.Println("...init GlobalLog...")
		log := logrus.New()
		log.SetLevel(logrus.InfoLevel)
		log.SetFormatter(&logrus.TextFormatter{
			// ForceColors:true,
			// EnvironmentOverrideColors:true,
			// FullTimestamp:true,
			TimestampFormat: "2006-01-02 15:04:05", //时间格式化
			// DisableLevelTruncation:true,
		})
		log.SetOutput(os.Stdout)
		logName := `log.txt`
		writer, err := rotatelogs.New(
			logName+".%Y%m%d%H%M%S",
			// WithLinkName为最新的日志建立软连接，以方便随着找到当前日志文件
			// rotatelogs.WithLinkName(logName),

			// WithRotationTime 设置日志分割的时间，这里设置为一小时分割一次
			rotatelogs.WithRotationTime(time.Hour*24*30), //文件旋转之间的间隔。默认情况下，日志每86400秒/一天旋转一次。注意:记住要利用时间。持续时间值。
			// WithMaxAge和WithRotationCount二者只能设置一个，
			// WithMaxAge设置文件清理前的最长保存时间，
			// WithRotationCount设置文件清理前最多保存的个数。 默认情况下，此选项是禁用的。
			// rotatelogs.WithMaxAge(time.Second*30),//默认每7天清除下日志文件
			rotatelogs.WithMaxAge(-1), //需要手动禁用禁用  默认情况下不清除日志，
			// rotatelogs.WithRotationCount(2),//清除除最新2个文件之外的日志，默认禁用
		)
		if err != nil {
			panic("config local file system for logger error: " + err.Error())
		}

		lfsHook := lfshook.NewHook(lfshook.WriterMap{
			logrus.DebugLevel: writer,
			logrus.InfoLevel:  writer,
			logrus.WarnLevel:  writer,
			logrus.ErrorLevel: writer,
			logrus.FatalLevel: writer,
			logrus.PanicLevel: writer,
		}, &logrus.TextFormatter{DisableColors: true})

		log.AddHook(lfsHook)

		// 初始化一些公共参数
		GlobalLog = log.WithFields(logrus.Fields{
			"app": "face",
		})
	}
	return GlobalLog
}

// UseLog().Error(map[string]interface{}{"request_id": request_id, "err": err})
