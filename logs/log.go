package logs

import (
	"fmt"
	"os"

	// "println"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"time"
)

type Logger struct {
	Mutex           sync.Mutex
	Prefix          string
	Skip            int64
	LogFunCCallPath int
	LogPath         string
}

const (
	ERROR  = "Error"
	ERRORF = "Errorf"
	INFO   = "Info"
	INFOF  = "Infof"
	WARN   = "Warn"
	WARNF  = "Warnf"
)

func NewLog(prefix string) (L *Logger) {
	L = new(Logger)
	L.Prefix = prefix
	L.LogFunCCallPath = 2
	// L.LogPath =
	return
}

var NewLogConfig Logger

func (L *Logger) Error(i ...interface{}) {
	L.writeLog(ERROR, "", i)
}

func (L *Logger) Errorf(format string, arg ...interface{}) {
	L.writeLog(ERRORF, format, arg)
}

func (L *Logger) Info(i ...interface{}) {
	L.writeLog(INFO, "", i)
}

func (L *Logger) Infof(format string, arg ...interface{}) {
	L.writeLog(INFOF, format, arg)
}

func (L *Logger) Warn(i ...interface{}) {
	L.writeLog(WARN, "", i)
}

func (L *Logger) Warnf(format string, arg ...interface{}) {
	L.writeLog(WARNF, format, arg)
}

func (L *Logger) writeLog(level, msg string, arg ...interface{}) {
	_, file, line, ok := runtime.Caller(L.LogFunCCallPath)
	if !ok {
		file = "?????"
		line = 0
	}
	fileSlice := strings.Split(file, "/")
	fileName := fileSlice[len(fileSlice)-1]
	if len(arg) > 0 {
		for _, v := range arg {
			msg += fmt.Sprint(v)
		}
	}
	msg = "[" + fileName + ":" + strconv.Itoa(line) + "]" + msg
	nowTime := time.Now().Format("2006-01-02")
	L.writeLogToFile(nowTime, msg, level)
	return
}
func (L *Logger) writeLogToFile(when, msg, level string) {
	//L.LogPath = "/Users/diannao/go/src"
	fileSlice := []string{"current.log", when + ".log"}
	chaInt := 2
	ch := make(chan int, chaInt)
	wg := new(sync.WaitGroup)
	for _, v := range fileSlice {
		ch <- 1
		wg.Add(1)
		v = L.LogPath + "/" + v
		go func(v string) {
			f, _ := os.OpenFile(v, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
			fTime := time.Now().Format("2006-01-02 15:04:05")
			fContent := fTime + " " + level + ":" + msg + "\n"
			buf := []byte(fContent)
			f.Write(buf)
			f.Close()
			<-ch
			wg.Done()
		}(v)
	}
	wg.Wait()

}

var loggerFunc Logger

func Error(i ...interface{}) {
	loggerFunc.Error(i)
}

func Errorf(format string, arg ...interface{}) {
	loggerFunc.Errorf(format, arg)
}

func Info(i ...interface{}) {
	loggerFunc.Info(i)
}

func Infof(format string, arg ...interface{}) {
	loggerFunc.Infof(format, arg)
}

func Warn(i ...interface{}) {
	loggerFunc.Warn(i)
}

func Warnf(format string, arg ...interface{}) {
	loggerFunc.Warnf(format, arg)
}
