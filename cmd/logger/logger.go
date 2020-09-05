package logger

import (
	"fmt"
	"runtime"
	"strings"
	"unsafe"

	"go.uber.org/zap"
)

var Log *zap.Logger

func LogInitAndStart() error {
	initLog()
	err := startLog()
	if err != nil {
		fmt.Println("log init and start fail: ", err)
		return err
	}
	fmt.Println("log init and start finish")
	return nil
}

func initLog() {
	if runtime.GOOS == OSWindows {
		if LogPath == "" {
			LogPath = LogLocalPath
		}
	}
	if runtime.GOOS == OSLinux {
		if LogPath == "" {
			LogPath = LogDefaultPath
		}
	}
	LogSize = LogDefaultSize
	LogBackups = LogDefaultBackups
	LogDay = LogDefaultDay
	LogLevel = InfoLevel
	LogGrade = StdOut
}

func startLog() error {
	Log = CreateCore(LogLevel, LogGrade)
	if Log == nil {
		fmt.Println("start failed")
	}
	return nil
}

func Panic() {
	if r := recover(); r != nil {
		var buf [4096]byte
		n := runtime.Stack(buf[:], false)
		bString := byteToStringUnSafe(buf[:n])
		strArr := make([]string, 0)
		strArr = append(strArr, fmt.Sprint(r)+": ")
		strArr = append(strArr, strings.Split(bString, "\t")...)
		for _, val := range strArr {
			Log.Error(StrPanic, zap.String(StrPanic, val))
		}
		panic(r)
	}
}

func PanicStr(str *string) {
	if r := recover(); r != nil {
		var buf [4096]byte
		n := runtime.Stack(buf[:], false)
		bString := byteToStringUnSafe(buf[:n])
		strArr := make([]string, 0)
		strArr = append(strArr, fmt.Sprint(r)+": ")
		strArr = append(strArr, strings.Split(bString, "\t")...)
		for _, val := range strArr {
			Log.Error(StrPanic, zap.String(StrPanic, val))
		}
		*str = strings.Join(strArr, "------")
		panic(r)
	}
}

func Field(v ...interface{}) zap.Field {
	return zap.Any("", fmt.Sprint(v))
}

func byteToStringUnSafe(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

func byteToString(b []byte) string {
	return fmt.Sprint("%s\n", string(b))
}

// import (
// 	"fmt"
// 	"log"
// 	"os"
// 	"path"
// 	"time"
//
// 	"dsight/common"
// )
//
// var debug *log.Logger
// var info *log.Logger
// var warn *log.Logger
// var error *log.Logger
//
// // var DOLPHIN_LOG string
// var LogPath = "/var/dana/log/eduplatform.log"
// var Days = 7
// var LogLevel = LOGINFO
//
// type EnumLog int
//
// const (
// 	LOGDEBUG EnumLog = iota // value 0
// 	LOGINFO                 // value 1
// 	LOGWARN                 // valeu 2
// 	LOGERRO                 // value 3
// )
//
// // 日志打印流
// func Initial() {
// 	// 读取当天的日期
// 	date := time.Now().Format("2006_01_02")
//
// 	logdir := fmt.Sprintf("%s_%s.log", common.LogPath, date)
// 	fmt.Println(logdir)
//
// 	os.MkdirAll(path.Dir(logdir), 0777)
//
// 	file, err := os.OpenFile(logdir, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)
// 	if err != nil {
// 		fmt.Println(err.Error())
// 	}
// 	LogLevel = EnumLog(common.LogLevel)
// 	Days = common.LogDays
//
// 	debug = log.New(file, "[	DEBUG	]", log.LstdFlags|log.Lshortfile)
// 	info = log.New(file, "[	INFO	]", log.LstdFlags|log.Lshortfile)
// 	warn = log.New(file, "[	WARN	]", log.LstdFlags|log.Lshortfile)
// 	error = log.New(file, "[	ERROR	]", log.LstdFlags|log.Lshortfile)
// }
//
// func Debug(v ...interface{}) {
// 	if LogLevel <= LOGDEBUG {
// 		debug.Println(v)
// 	}
// }
//
// func Info(v ...interface{}) {
// 	if LogLevel <= LOGINFO {
// 		info.Println(v)
// 	}
// }
//
// func Infof(format string, v ...interface{}) {
// 	if LogLevel <= LOGINFO {
// 		info.Printf(format, v...)
// 	}
// }
//
// func Warn(v ...interface{}) {
// 	if LogLevel <= LOGWARN {
// 		warn.Println(v)
// 	}
// }
//
// func Error(v ...interface{}) {
// 	if LogLevel <= LOGERRO {
// 		error.Println(v)
// 	}
// }
