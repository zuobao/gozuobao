package logger

import (
	"errors"
	"fmt"
	"github.com/Sirupsen/logrus"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"
)

var loggerMap = map[string]*logrus.Logger{}
var loggerFileMap = map[string]*os.File{}

// 控制台日志
var stdoutLogger *log.Logger
var stderrLogger *log.Logger

// 文件日志
var infoLogger *log.Logger
var debugLogger *log.Logger
var errorLogger *log.Logger

var infoLoggerFile *os.File
var debugLoggerFile *os.File
var errorLoggerFile *os.File

var DebugEnabled = false

func checkErr(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func init() {

	stdoutLogger = log.New(os.Stdout, "", log.LstdFlags|log.Lshortfile)
	stderrLogger = log.New(os.Stderr, "", log.LstdFlags|log.Lshortfile)

	//	stderrLogger.

	//	stderrLogger.Out = os.Stderr

	//	loggerMap["stdout"] = stdoutLogger
	//	loggerMap["stderr"] = stderrLogger

	if err := checkDir(LogFileRoot); err != nil {
		stderrLogger.Panic("open log directory error: " + err.Error())
	}
	if err := checkDir(LogEventFileRoot); err != nil {
		stderrLogger.Panic("open log directory error: " + err.Error())
	}

	var err error
	infoLogger, infoLoggerFile, err = createLogger("info", "", log.LstdFlags|log.Lshortfile)
	checkErr(err)

	errorLogger, errorLoggerFile, err = createLogger("error", "", log.LstdFlags|log.Lshortfile)
	checkErr(err)

	debugLogger, debugLoggerFile, err = createLogger("debug", "", log.LstdFlags|log.Lshortfile)
	checkErr(err)

	//	stdoutLogger.Level = logrus.Debug
	//	debugLogger.Level = logrus.Debug

	//	jsonFormatter := new (logrus.JSONFormatter)

	//	stdoutLogger.Formatter = jsonFormatter
	//	stderrLogger.Formatter = jsonFormatter
}

// 日志文件根目录，默认为当前目录下的logs目录
var LogFileRoot = "./logs"
var LogEventFileRoot = "./logs/events"
var LogFileSuffix = ".log"

// 确保目录存在
func checkDir(dir string) error {
	stat, err := os.Stat(dir)
	if os.IsNotExist(err) {
		err = os.MkdirAll(dir, os.FileMode(0777))
		stat, err = os.Stat(dir)
	}
	if err != nil {
		return err
	}

	if !stat.IsDir() {
		return errors.New("not a directory: " + dir)
	}
	return nil
}

func createLoggerFile(filename string) (*os.File, error) {

	fullpath := LogFileRoot + "/" + filename + LogFileSuffix
	fullpath = filepath.Clean(fullpath)

	var f *os.File
	var err error
	f, err = os.OpenFile(fullpath, os.O_APPEND|os.O_WRONLY|os.O_CREATE, os.FileMode(0777))
	if err != nil {
		return nil, err
	}

	return f, nil
}

func createLogrusLogger(filename string) (*logrus.Logger, *os.File, error) {
	f, err := createLoggerFile("events/" + filename)
	if err != nil {
		return nil, nil, err
	}

	l := logrus.New()
	l.Out = f
	l.Formatter = new(logrus.JSONFormatter)

	return l, f, nil
}

func createLogger(filename string, prefix string, flags int) (*log.Logger, *os.File, error) {
	f, err := createLoggerFile(filename)
	if err != nil {
		return nil, nil, err
	}

	l := log.New(f, prefix, flags)

	return l, f, nil
}

//type Logger logrus.Logger

//func GetStdout() *log.Logger {
//	return stdoutLogger
//}

func Fatalf(format string, args ...interface{}) {
	s := fmt.Sprintf(format, args...)
	stderrLogger.Output(2, s)
	infoLogger.Output(2, s)
	errorLogger.Output(2, s)
	debugLogger.Output(2, s)

	os.Exit(1)
}

func Fatalln(args ...interface{}) {
	s := fmt.Sprintln(args...)
	stderrLogger.Output(2, s)
	infoLogger.Output(2, s)
	errorLogger.Output(2, s)
	debugLogger.Output(2, s)

	os.Exit(1)
}

func Infoln(args ...interface{}) {
	s := fmt.Sprintln(args...)
	stdoutLogger.Output(2, s)
	infoLogger.Output(2, s)

	debugLogger.Output(2, s)

}

func Infof(format string, args ...interface{}) {
	s := fmt.Sprintf(format, args...)
	stdoutLogger.Output(2, s)
	infoLogger.Output(2, s)
	debugLogger.Output(2, s)

}

func Errorln(args ...interface{}) {
	s := fmt.Sprintln(args...)
	//	stdoutLogger.Output(2, s)
	stderrLogger.Output(2, s)
	errorLogger.Output(2, s)
	debugLogger.Output(2, s)

}

func ErrorlnWithDepth(depth int, args ...interface{}) {
	s := fmt.Sprintln(args...)
	//	stdoutLogger.Output(2, s)
	stderrLogger.Output(depth, s)
	errorLogger.Output(depth, s)
	debugLogger.Output(depth, s)

}

func Errorf(format string, args ...interface{}) {
	s := fmt.Sprintf(format, args...)
	//	stdoutLogger.Output(2, s)
	stderrLogger.Output(2, s)
	errorLogger.Output(2, s)

	debugLogger.Output(2, s)

}

func Warnf(format string, args ...interface{}) {
	s := fmt.Sprintf(format, args...)
	stderrLogger.Output(2, s)
	errorLogger.Output(2, s)
	debugLogger.Output(2, s)
}

func Warnln(args ...interface{}) {
	s := fmt.Sprintln(args...)
	stderrLogger.Output(2, s)
	errorLogger.Output(2, s)
	debugLogger.Output(2, s)
}

func Debugln(args ...interface{}) {
	if DebugEnabled {
		s := fmt.Sprintln(args...)
		stdoutLogger.Output(2, s)
		debugLogger.Output(2, s)
	}
}

func Debugf(format string, args ...interface{}) {
	if DebugEnabled {
		s := fmt.Sprintf(format, args...)
		stdoutLogger.Output(2, s)
		debugLogger.Output(2, s)
	}
}

func Get(filename string) (*logrus.Logger, error) {
	filename = strings.ToLower(filename)

	var l *logrus.Logger
	var f *os.File
	var err error
	var ok bool

	l, ok = loggerMap[filename]
	if !ok {
		l, f, err = createLogrusLogger(filename)
		if err != nil {
			return nil, err
		}
		loggerMap[filename] = l
		loggerFileMap[filename] = f
	}

	return l, nil
}

type LogHttpRequest struct {
	Ip, UserAgent string
	Method        string
	Path          string
	QueryString   string
	Start, End    time.Time
}
