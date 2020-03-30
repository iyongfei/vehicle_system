package logger

import (
	"errors"
	"fmt"
	"os"
	"path"
	"runtime"
	"sync"
	"time"
)

const (
	LOG_LEVEL_INFO = iota
)

const (
	defaultDepth = 3
)

var (
	maxFileCap = 1024 * 1024 * 50
	logDir = "vlog"
)

var (
	OsGetWdFail = errors.New("os getwd msg failed")
	ArgsInvaild      = errors.New("args can be vaild")
	ObtainFileFail   = errors.New("obtain file failed")
	OpenFileFail     = errors.New("open file failed")
	GetLineNumFail   = errors.New("get line num faild")
	WriteLogInfoFail = errors.New("write log msg failed")
	OStatPathFail = errors.New("os stat path failed")
)

type VLogger struct {
	m_FileDir       string
	m_FileName      string
	m_FileHandle    *os.File
	m_Level         int
	m_Depth         int
	m_nexDay        time.Time
	m_MaxLogDataNum int
	m_mu            sync.Mutex
}

func defaultNew() *VLogger {
	return &VLogger{
		m_FileDir:       "",
		m_FileName:      "",
		m_FileHandle:    nil,
		m_Level:         0,
		m_Depth:         defaultDepth,
		m_MaxLogDataNum: maxFileCap,
	}
}

func NewRealStLogger() (*VLogger,error) {
	//获取log文件夹的路径
	logDirPwd, err := os.Getwd()
	if err != nil {
		fmt.Println(OsGetWdFail)
		os.Exit(1)
		return nil,OsGetWdFail
	}

	logger := defaultNew()
	logger.m_FileDir = logDirPwd+"/"+logDir

	err = logger.obtainLofFile()

	if err != nil {
		return nil,ObtainFileFail
	}
	return logger,nil
}

func (this *VLogger) obtainLofFile() error {
	fileDir := this.m_FileDir
	//文件夹为空
	if fileDir == "" {
		fmt.Println(ArgsInvaild)
		os.Exit(1)
		return ArgsInvaild
	}

	//时间文件夹log20181125
	//destFilePath := fmt.Sprintf("%s", "wxpay_log")
	flag, err := IsExist(fileDir)
	if err != nil {
		return OStatPathFail
	}
	if !flag {
		os.MkdirAll(fileDir, os.ModePerm)
	}
	//文件夹存在,直接以创建的方式打开文件
	destFilePath := fileDir + "/" // 格式为log20181125/log_1_20181125.txt
	logFilePath := fmt.Sprintf("%s%d%d%d%s", destFilePath, time.Now().Year(), time.Now().Month(),
		time.Now().Day(), ".log")

	fileHandle, err := os.OpenFile(logFilePath, os.O_CREATE|os.O_APPEND|os.O_RDWR, os.ModePerm)
	if err != nil {
		fmt.Println(OpenFileFail, err.Error())
	}

	this.m_FileHandle = fileHandle
	this.m_FileName = logFilePath
	//设置下次创建文件的时间
	nextDay := time.Unix(time.Now().Unix()+(24 * 3600), 0)
	nextDay = time.Date(nextDay.Year(), nextDay.Month(), nextDay.Day(), 0, 0, 0,
		0, nextDay.Location())
	this.m_nexDay = nextDay
	return nil
}

///index
func (this *VLogger) FormatWriteLogMsg( logMsg string) {
	this.m_mu.Lock()
	defer this.m_mu.Unlock()
	now := time.Now()

	if now.Unix() > this.m_nexDay.Unix() /**|| int(fileSize) > this.m_MaxLogDataNum*/ {
		err := this.obtainLofFile()
		if err != nil {
			fmt.Println(ObtainFileFail)
		}
	}

	flag := GetLoggerLevel(0)
	_, file, line, ok := runtime.Caller(this.m_Depth)
	if ok == false {
		fmt.Println(GetLineNumFail)
	}
	name := path.Base(file)
	timer := time.Now().Format("2006-01-02 15:04:05.000")

	fmt.Println(this.m_FileHandle.Name())//2018-11-25 17:55:49.845 [INFO]: [context.go:108] INFO 12
	_, err := Write(this.m_FileHandle, fmt.Sprintf("%s %s [%s:%d] %s\n", timer, flag, name, line, logMsg))
	if err != nil {
		fmt.Println(WriteLogInfoFail, err.Error())
	}
}


func (this *VLogger) INFO(format string, args ...interface{}) {
	this.FormatWriteLogMsg(fmt.Sprintf(format, args...))
}


func GetLoggerLevel(level int) string {
	switch level {
	case LOG_LEVEL_INFO:
		return "[INFO]:"
	default:
		return ""
	}
}
