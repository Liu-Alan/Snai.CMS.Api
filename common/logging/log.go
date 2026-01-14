package logging

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/fatih/color"
)

var logger *Logger //系统日志
var _logLevels = []string{"", "Debug", "Info", "Warn", "Error", "Fatal"}

// LogLevel 日志等级
type LogLevel int

// 日志等级
const (
	LogDebug LogLevel = 1
	LogInfo  LogLevel = 2
	LogWarn  LogLevel = 3
	LogError LogLevel = 4
	LogFatal LogLevel = 5
)

// Logger 日志
type Logger struct {
	level     LogLevel
	toFile    bool
	toConsole bool
	folder    string
	file      *os.File
	fileDate  string
}

// InitLog 初始化系统日志
func InitLog(level LogLevel, targets int, folder string) {
	logger = &Logger{level: level, folder: folder}
	if targets&1 == 1 {
		logger.toConsole = true
	}
	if targets&2 == 2 {
		logger.toFile = true
	}
	logger.init()
}

// Log 无视级别输出
func Log(format string, v ...interface{}) {
	txt := logger.getLogTextInner(format, v...)
	txt = fmt.Sprintf("%v %v", time.Now().Format("15:04:05"), txt)

	if logger.toConsole {
		fmt.Print(txt)
	}
	if logger.toFile {
		logger.getFile().Write([]byte(txt))
	}
}

// Debug output logs at debug level
func Debug(format string, v ...interface{}) {
	if logger.level > LogDebug {
		return
	}

	txt := logger.getLogText(1, format, v...)
	if logger.toConsole {
		color.Cyan(txt)
	}
	if logger.toFile {
		logger.getFile().Write([]byte(txt))
	}
}

// Info output logs at info level
func Info(format string, v ...interface{}) {
	if logger.level > LogInfo {
		return
	}

	txt := logger.getLogText(2, format, v...)
	if logger.toConsole {
		color.Green(txt)
	}
	if logger.toFile {
		logger.getFile().Write([]byte(txt))
	}
}

// Warn output logs at warn level
func Warn(format string, v ...interface{}) {
	if logger.level > LogWarn {
		return
	}

	txt := logger.getLogText(3, format, v...)
	if logger.toConsole {
		color.Yellow(txt)
	}
	if logger.toFile {
		logger.getFile().Write([]byte(txt))
	}
}

// Error output logs at error level
func Error(format string, v ...interface{}) {
	if logger.level > LogError {
		return
	}

	txt := logger.getLogText(4, format, v...)
	if logger.toConsole {
		color.Red(txt)
	}
	if logger.toFile {
		logger.getFile().Write([]byte(txt))
	}
}

// Fatal output logs at fatal level
func Fatal(format string, v ...interface{}) {
	if logger.level > LogFatal {
		return
	}
	txt := logger.getLogText(5, format, v...)
	if logger.toConsole {
		c := color.New(color.BgRed, color.FgWhite)
		c.Print(txt)
	}
	if logger.toFile {
		logger.getFile().Write([]byte(txt))
	}
	os.Exit(1)
}

// 获取文件
func (x *Logger) getFile() *os.File {
	now := time.Now().Format("2006-01-02")
	if now != x.fileDate {
		x.init()
	}
	return x.file
}

// 获取日志文本信息，不带日期和类型
func (x *Logger) getLogTextInner(format string, v ...interface{}) string {
	txt := ""

	if len(v) == 0 {
		txt = format
	} else {
		txt = fmt.Sprintf(format, v...)
	}

	if !strings.HasSuffix(txt, "\n") {
		txt += "\n"
	}
	return txt
}

func (x *Logger) getLogText(level int, format string, v ...interface{}) string {
	txt := x.getLogTextInner(format, v...)
	return fmt.Sprintf("%v [%v] %v", time.Now().Format("2006-01-02 15:04:05"), _logLevels[level], txt)
}

func (x *Logger) init() {
	outputs := []string{}
	x.fileDate = time.Now().Format("2006-01-02")

	if x.toConsole {
		outputs = append(outputs, "Stdout")
	}

	if x.toFile {
		outputs = append(outputs, "File")

		if x.file != nil {
			x.file.Close()
		}

		filePath := time.Now().Format("2006-01-02") + ".log"
		folder := x.folder

		// 创建文件夹
		if _, err := os.Stat(folder); os.IsNotExist(err) {
			os.Mkdir(folder, os.ModePerm)
		}

		f, _ := os.OpenFile(folder+"/"+filePath, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
		x.file = f
	}

	txt := fmt.Sprintf("日志级别:%v, 输出位置:%v\n", _logLevels[x.level], strings.Join(outputs, ", "))
	c := color.New(color.Bold)

	c.Printf(txt)
	if x.toFile {
		x.file.Write([]byte(txt))
	}
}

// 用于sql日志
var slogger *Logger // sql日志
type SlogWriter struct {
}

// InitSlog 初始化sql日志
func InitSlog(level LogLevel, targets int, folder string) {
	slogger = &Logger{level: level, folder: folder}
	if targets&1 == 1 {
		slogger.toConsole = true
	}
	if targets&2 == 2 {
		slogger.toFile = true
	}
	slogger.init()
}

func (w SlogWriter) Printf(format string, args ...interface{}) {
	txt := slogger.getLogText(2, format, args...)
	if slogger.toConsole {
		color.Yellow(txt)
	}
	if slogger.toFile {
		slogger.getFile().Write([]byte(txt))
	}
}
