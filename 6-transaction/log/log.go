package log

import (
	"io/ioutil"
	"log"
	"os"
	"sync"
)

/**
[info ] 颜色为蓝色，[error] 为红色。
使用 log.Lshortfile 支持显示文件名和代码行号。
暴露 Error，Errorf，Info，Infof 4个方法。
*/

var (
	errorLog = log.New(os.Stdout, "\033[31m[error]\033[0m ", log.LstdFlags|log.Lshortfile)
	infoLog  = log.New(os.Stdout, "\033[34m[info ]\033[0m ", log.LstdFlags|log.Lshortfile)
	loggers  = []*log.Logger{errorLog, infoLog}
	mu       sync.Mutex
)

// 定义log methods

var (
	Error  = errorLog.Println
	Errorf = errorLog.Printf
	Info   = infoLog.Println
	Infof  = infoLog.Printf
)

/**
设置日志的层级(InfoLevel, ErrorLevel, Disabled)。
*/

// log level
const (
	InfoLevel = iota
	ErrorLevel
	Disabled
)

// SetLevel
//  @Description: 设置loglevel等级
//  @param level：0，1，2
//
func SetLevel(level int) {
	mu.Lock()
	defer mu.Unlock()

	for _, loggers := range loggers {
		loggers.SetOutput(os.Stdout)
	}

	if ErrorLevel < level {
		errorLog.SetOutput(ioutil.Discard)
	}

	if InfoLevel < level {
		infoLog.SetOutput(ioutil.Discard)
	}
}
