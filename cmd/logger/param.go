package logger

var (
	// 日志属性
	LogSize    int    // megabytes
	LogBackups int    // 备份数量
	LogDay     int    // 保存天数
	LogPath    string // 日志路径
	LogLevel   int    // 日志级别
	LogGrade   string // 日志使用等级
)

const (
	StrError = "error: "
	StrPanic = "panic: "
	StrWarn  = "warn: "
	StrInfo  = "info: "
	StrDebug = "debuf: "
)

// 日志默认参数
const (
	// 日志名称
	LogName = "daily.log"
	// 日志本地路径
	LogLocalPath = "public/logger/log"
	// 日志默认linux路径
	LogDefaultPath = "/root/daily/.log"

	// 日志时间精确度（秒）
	TimeFormat = "2006-01-02 15:04:05"

	// 系统类型
	OSWindows = "windows"
	OSLinux   = "linux"

	// 日志属性
	LogDefaultSize    = 10 // megabytes
	LogDefaultBackups = 3  // 备份数量
	LogDefaultDay     = 7  // 保存天数

	// 日志级别
	DebugLevel = 1
	InfoLevel  = 2
	WarnLevel  = 3
	ErrorLevel = 4
	PanicLevel = 5

	// 使用级别
	File    = "file"
	StdOut  = "stdout"
	Analyze = "analyze"
)
