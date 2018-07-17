package logger

import (
	"os"
	"time"

	"github.com/evalphobia/logrus_sentry"
	raven "github.com/getsentry/raven-go"
	"github.com/sirupsen/logrus"
)

// Log Create a new instance of the logger. You can have any number of instances.
var log *logrus.Logger

// Init 初始化
func Init() {
	log = logrus.New()
	log.Out = os.Stdout
	// 只记录警告严重程度或以上。
	log.SetLevel(logrus.InfoLevel)
}

// Fatalf 致命
func Fatalf(format string, args ...interface{}) {
	log.Fatalf(format, args)
}

// Warningf 警告
func Warningf(format string, args ...interface{}) {
	log.Warningf(format, args)
}

// Errorf 错误
func Errorf(format string, args ...interface{}) {
	log.Errorf(format, args)
}

// Infof 信息
func Infof(format string, args ...interface{}) {
	log.Infof(format, args)
}

// Printf 打印
func Printf(format string, args ...interface{}) {
	log.Printf(format, args)
}

// Debugln 测试
func Debugln(args ...interface{}) {
	log.Debugln(args)
}

// Infoln 消息
func Infoln(args ...interface{}) {
	log.Infoln(args)
}

// Println 打印
func Println(args ...interface{}) {
	log.Println(args)
}

// Warnln 警告
func Warnln(args ...interface{}) {
	log.Warnln(args)
}

// Warningln 警告
func Warningln(args ...interface{}) {
	log.Warningln(args)
}

// Errorln 错误
func Errorln(args ...interface{}) {
	log.Errorln(args)
}

// Fatalln 严重
func Fatalln(args ...interface{}) {
	log.Fatalln(args)
}

// Panicln 恐慌
func Panicln(args ...interface{}) {
	log.Panicln(args)
}

// RegisterSentry register sentry hook with a DSN key
func RegisterSentry(DSN, serviceName string) {
	tags := map[string]string{
		// "site":   env + "-api.dmg.com",
		"server": serviceName,
	}

	client, err := raven.NewClient(DSN, tags)
	if err != nil {
		log.Fatal(err)
	}

	hook, err := logrus_sentry.NewWithClientSentryHook(client, []logrus.Level{
		logrus.PanicLevel,
		logrus.FatalLevel,
		logrus.ErrorLevel,
		logrus.WarnLevel,
	})
	hook.Timeout = 2 * time.Second

	if err != nil {
		log.Error("Logger RegisterSentry Failed:", err.Error())
		return
	}
	log.Hooks.Add(hook)
}
