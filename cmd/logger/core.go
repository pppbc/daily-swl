package logger

import (
	"os"
	"path/filepath"
	"runtime"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

func CreateCore(level int, grade string) *zap.Logger {
	if grade == StdOut {
		return NewDefaulCoreV2(level)
	}
	if grade == Analyze {
		// null function
		return NewDefaulCoreV3(level)
	} else {
		// defaul use NewDefaulCoreV1
		return NewDefaulCoreV1(level)
	}
}

func NewEncoderConfig() zapcore.EncoderConfig {
	return zapcore.EncoderConfig{
		// Keys can be anything except the empty string.
		TimeKey:        "T",
		LevelKey:       "L",
		NameKey:        "N",
		CallerKey:      "C",
		MessageKey:     "M",
		StacktraceKey:  "S",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.CapitalLevelEncoder,
		EncodeTime:     TimeEncoder,
		EncodeDuration: zapcore.StringDurationEncoder,
		EncodeCaller:   zapcore.FullCallerEncoder,
	}
}

func TimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format(TimeFormat))
}

// only log file
func NewDefaulCoreV1(level int) *zap.Logger {
	fileName := filepath.Join(LogPath, LogName)
	w := zapcore.AddSync(&lumberjack.Logger{
		Filename:   fileName,
		MaxSize:    LogSize, // megabytes
		MaxBackups: LogBackups,
		MaxAge:     LogDay, // days
	})
	core := NewZapCore(level, w)

	Logger := zap.New(core, zap.AddCaller())
	return Logger
}

// log file and stdout
func NewDefaulCoreV2(level int) *zap.Logger {
	if runtime.GOOS == OSWindows {
		LogPath = LogLocalPath
	}
	if runtime.GOOS == OSLinux {
		if LogPath == "" {
			LogPath = LogDefaultPath
		}
	}
	fileName := filepath.Join(LogPath, LogName)
	w := zapcore.AddSync(&lumberjack.Logger{
		Filename:   fileName,
		MaxSize:    LogSize, // megabytes
		MaxBackups: LogBackups,
		MaxAge:     LogDay, // days
	})
	core := NewZapCore(level, w, zapcore.AddSync(os.Stdout))

	Logger := zap.New(core, zap.AddCaller())
	return Logger
}

// log file、stdout and analyze
// the function is null in now
func NewDefaulCoreV3(level int) *zap.Logger {
	return nil
}

// create a zapcore
func NewZapCore(level int, w ...zapcore.WriteSyncer) zapcore.Core {
	var loglevel zapcore.Level
	switch level {
	case 1:
		loglevel = zap.DebugLevel
	case 2:
		loglevel = zap.InfoLevel
	case 3:
		loglevel = zap.WarnLevel
	case 4:
		loglevel = zap.ErrorLevel
	case 5:
		loglevel = zap.PanicLevel
	default:
		loglevel = zap.DebugLevel
	}
	return zapcore.NewCore(
		zapcore.NewConsoleEncoder(NewEncoderConfig()),
		zapcore.NewMultiWriteSyncer(w...),
		loglevel)
}

// func testhhh(){
// 	// First, define our level-handling logic.
// 	// 仅打印Error级别以上的日志
// 	highPriority := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
// 		return lvl >= zapcore.ErrorLevel
// 	})
// 	// 打印所有级别的日志
// 	lowPriority := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
// 		return lvl >= zapcore.DebugLevel
// 	})
//
// 	hook := lumberjack.Logger{
// 		Filename:   "/tmp/abc.log",
// 		MaxSize:    1024, // megabytes
// 		MaxBackups: 3,
// 		MaxAge:     7,    //days
// 		Compress:   true, // disabled by default
// 	}
//
//
// 	topicErrors := zapcore.AddSync(ioutil.Discard)
// 	fileWriter := zapcore.AddSync(&hook)
//
// 	// High-priority output should also go to standard error, and low-priority
// 	// output should also go to standard out.
// 	consoleDebugging := zapcore.Lock(os.Stdout)
//
// 	// Optimize the Kafka output for machine consumption and the console output
// 	// for human operators.
// 	kafkaEncoder := zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig())
// 	consoleEncoder := zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig())
//
// 	// Join the outputs, encoders, and level-handling functions into
// 	// zapcore.Cores, then tee the four cores together.
// 	core := zapcore.NewTee(
// 		// 打印在kafka topic中（伪造的case）
// 		zapcore.NewCore(kafkaEncoder, topicErrors, highPriority),
// 		// 打印在控制台
// 		zapcore.NewCore(consoleEncoder, consoleDebugging, lowPriority),
// 		// 打印在文件中
// 		zapcore.NewCore(consoleEncoder, fileWriter, highPriority),
// 	)
//
// 	// From a zapcore.Core, it's easy to construct a Logger.
// 	logger := zap.New(core)
// 	defer logger.Sync()
// 	logger.Info("constructed a info logger", zap.Int("test", 1))
// 	logger.Log.Error(logger.StrError,logger.Field(err))("constructed a error logger", zap.Int("test", 2))
// }
