package logger

import (
	"errors"
	"fmt"
	"os"
	"path"
	"runtime"
	"strings"
	"sync"
	"time"
)

const (
	LOG_LEVEL_PRINT = 0
	LOG_LEVEL_INFO = 1
	LOG_LEVEL_ERROR = 2
	DEFAULT_DEPTH = 3
	DEL_LOG_DAYS = 7
	DAY_SECONDS = 24*3600
)

var (
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
	m_Depth         int
	m_nexDay        time.Time
	m_DelDay    uint
	m_mu            sync.Mutex
}

func defaultNew() *VLogger {
	return &VLogger{
		m_FileDir:       "",
		m_FileName:      "",
		m_FileHandle:    nil,
		m_Depth:         DEFAULT_DEPTH,
		m_DelDay:DEL_LOG_DAYS,
	}
}
func NewVLogger() (*VLogger,error) {
	logDirPwd, err := os.Getwd()
	fmt.Println(logDirPwd,"logDirPwd")
	if err != nil {
		fmt.Println(OsGetWdFail,err.Error())
		os.Exit(1)
		return nil,OsGetWdFail
	}

	logger := defaultNew()
	logger.m_FileDir = logDirPwd+"/"+logDir

	err = logger.obtainLofFile()

	if err != nil {
		fmt.Println(ObtainFileFail,err.Error())
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
	fileDirExist, err := IsExist(fileDir)
	if err != nil {
		fmt.Println(OStatPathFail,err.Error())
		return OStatPathFail
	}
	if !fileDirExist {
		os.MkdirAll(fileDir, os.ModePerm)
	}
	//文件夹存在,直接以创建的方式打开文件
	destFilePath := fileDir + "/" // 格式为log20181125/log_1_20181125.txt
	logFilePath := fmt.Sprintf("%s%d%d%d%s", destFilePath, time.Now().Year(), time.Now().Month(),
		time.Now().Day(), ".log")

	fileHandle, err := os.OpenFile(logFilePath, os.O_CREATE|os.O_APPEND|os.O_RDWR, os.ModePerm)
	if err != nil {
		fmt.Println(OpenFileFail,err.Error())
		return OpenFileFail
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

func (this *VLogger) FormatWriteLogMsg(level int, logMsg string) {
	this.m_mu.Lock()
	defer this.m_mu.Unlock()
	now := time.Now()

	if now.Unix() > this.m_nexDay.Unix() /**|| int(fileSize) > this.m_MaxLogDataNum*/ {
		err := this.obtainLofFile()
		if err != nil {
			fmt.Println(ObtainFileFail,err.Error())
		}
	}

	for i := 0;i < DEL_LOG_DAYS ; i++ {
		this.RemoveTimeOutLogFolder(this.m_DelDay+uint(i))
	}

	logLevel := GetLoggerLevel(level)
	_, file, line, ok := runtime.Caller(this.m_Depth)
	if ok == false {
		fmt.Println(GetLineNumFail)
	}
	pathBase := path.Base(file)
	timer := time.Now().Format("2006-01-02 15:04:05")
	_, err := Write(this.m_FileHandle, fmt.Sprintf("%s [%s:%d] %s %s\n", timer,pathBase, line, logLevel,logMsg))
	if err != nil {
		fmt.Println(WriteLogInfoFail, err.Error())
	}
}

func (this *VLogger)RemoveTimeOutLogFolder(uiDayAgo uint)  {
	timeNow := time.Now().Unix()
	timeAgo := timeNow - int64(uiDayAgo*DAY_SECONDS)
	t := time.Unix(timeAgo,0)

	destFilePath := this.m_FileDir + "/"

	folderName := fmt.Sprintf("%s%d%d%d",destFilePath,t.Year(),t.Month(),t.Day())
	os.RemoveAll(folderName)
}


func (this *VLogger) Print(format string, args ...interface{}) {
	lineFeed:=strings.HasSuffix(format,"\n")
	if !lineFeed{
		format = format + "\n"
	}
	fmt.Printf(format,args...)
}

func (this *VLogger) Info(format string, args ...interface{}) {
	this.FormatWriteLogMsg(LOG_LEVEL_INFO,fmt.Sprintf(format, args...))
}

func (this *VLogger) Error(format string, args ...interface{}) {
	this.FormatWriteLogMsg(LOG_LEVEL_ERROR,fmt.Sprintf(format, args...))
}

func GetLoggerLevel(level int) string {
	switch level {
	case LOG_LEVEL_PRINT:
		return "[PRINT]:"
	case LOG_LEVEL_INFO:
		return "[INFO]:"
	case LOG_LEVEL_ERROR:
		return "[ERROR]:"
	default:
		return ""
	}
}
