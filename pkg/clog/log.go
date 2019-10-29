/***
文本日志
*/
package clog

import (
	"bufio"
	"fmt"
	"os"
	"path"
	"time"

	"github.com/seerx/got"
	"github.com/sirupsen/logrus"

	roratelogs "github.com/lestrrat/go-file-rotatelogs"
	"github.com/rifflock/lfshook"
)

var Log = logrus.New()

// InitCLog 初始化日志
func InitCLog(cfg *LogConfigure) {
	fmt.Println("日志", cfg.Path)
	//tf = got.NewTimeFormatter()
	Log.Out = os.Stdout
	// Log.Formatter = &myFormatter{}
	// logrus.TextFormatter{}
	Log.Formatter = &logrus.JSONFormatter{}

	_, err := os.Lstat(cfg.Path)
	if err != nil {
		if os.IsNotExist(err) {
			// 创建路径
			got.MakeDirecotories(cfg.Path)
		} else {
			fmt.Println("err", err)
		}
	}

	baseLogFile := path.Join(cfg.Path, "log")

	writer, err := roratelogs.New(
		baseLogFile+".%Y%m%d%H%M",
		roratelogs.WithLinkName(baseLogFile),
		roratelogs.WithMaxAge(30*24*time.Hour),
		roratelogs.WithRotationTime(24*time.Hour),
	)
	if err != nil {
		fmt.Println("err", err)
	}

	switch cfg.Level {
	case "debug":
		Log.SetLevel(logrus.DebugLevel)
		Log.SetOutput(os.Stderr)
	case "info":
		Log.SetLevel(logrus.InfoLevel)
		setNull()
	case "warn":
		Log.SetLevel(logrus.WarnLevel)
		setNull()
	case "error":
		Log.SetLevel(logrus.ErrorLevel)
		setNull()
	default:
		Log.SetLevel(logrus.InfoLevel)
		setNull()
	}
	//logrus.DebugLevel
	lfHook := lfshook.NewHook(lfshook.WriterMap{
		logrus.DebugLevel: writer, // 为不同级别设置不同的输出目的
		logrus.InfoLevel:  writer,
		logrus.WarnLevel:  writer,
		logrus.ErrorLevel: writer,
		logrus.FatalLevel: writer,
		logrus.PanicLevel: writer,
	}, &logrus.JSONFormatter{})
	Log.AddHook(lfHook)

}

func setNull() {
	src, err := os.OpenFile(os.DevNull, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		fmt.Println("err", err)
	}
	writer := bufio.NewWriter(src)
	Log.SetOutput(writer)
}
