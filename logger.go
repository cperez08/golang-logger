package gologger

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"sync"
	"time"
)

const (

	//LOG LEVELS

	// TRACE Indicates the lower log level
	TRACE = 1 << iota
	// DEBUG Indicates the low log level
	DEBUG
	// INFO Indicates the mid log level
	INFO
	// WARN Indicates the hig log level
	WARN
	// ERROR Indicates the higer log level
	ERROR

	// LOG MODES

	//CUSTOM you define based on flag paramenter which levels activate
	CUSTOM = "CUSTOM"

	//PROD autodefine the  best configuration for production
	PROD = "PRODUCTION"

	//DEV autodefine the  best configuration for development env
	DEV = "DEVELOPMENT"
)

//Logger ...
type Logger struct {

	//Mutext ...
	lock sync.Mutex

	//Trace ...
	Trace *log.Logger

	//Debug ...
	Debug *log.Logger

	//Info ...
	Info *log.Logger

	//Warning ...
	Warning *log.Logger

	//Error ...
	Error *log.Logger

	//File log
	FileLog *os.File

	//Indicates if the rotation is enabled, applies if the library manages the file
	Rotation bool

	// Value in KB for the log rotation, applies if the librar manage the file
	RotationWeight uint
}

//NewLogger ...
func NewLogger(file string, logMode string, flags int) *Logger {

	useFile := handleLogFile(file)

	if useFile == nil {

		return nil
	}

	loggerRs := &Logger{

		FileLog:        useFile,
		Rotation:       false,
		RotationWeight: 0,
	}

	switch logMode {

	case DEV:

		loggerRs.initDevelopmentLoggers()

	case PROD:

		loggerRs.initProductionLoggers()

	default:

		loggerRs.intiCustomLoggers(flags)

	}

	return loggerRs

}

//NewLoggerWithOpenFile ...
func NewLoggerWithOpenFile(openFile *os.File, logMode string, flags int) *Logger {

	if openFile == nil {

		return nil
	}

	loggerRs := &Logger{

		FileLog:        openFile,
		Rotation:       false,
		RotationWeight: 0,
	}

	switch logMode {

	case DEV:

		loggerRs.initDevelopmentLoggers()

	case PROD:

		loggerRs.initProductionLoggers()

	default:

		loggerRs.intiCustomLoggers(flags)

	}

	return loggerRs

}

//NewLoggerWithRotation ...
func NewLoggerWithRotation(file string, logMode string, flags int, maxLogFile uint) *Logger {

	useFile := handleLogFile(file)

	if useFile == nil {

		return nil
	}

	loggerRs := &Logger{
		FileLog:        useFile,
		Rotation:       true,
		RotationWeight: maxLogFile,
	}

	switch logMode {

	case DEV:

		loggerRs.initDevelopmentLoggers()

	case PROD:

		loggerRs.initProductionLoggers()

	default:

		loggerRs.intiCustomLoggers(flags)

	}

	return loggerRs

}

//CloseFile manual closing file
func (l *Logger) CloseFile() bool {

	if err := l.FileLog.Close(); err != nil {

		fmt.Print("error closing log file")
		return false
	}

	return true
}

//RotateLog ...
func (l *Logger) RotateLog() bool {

	var errCreatingFile error
	filename := filepath.Base(l.FileLog.Name())
	filePath := filepath.Dir(l.FileLog.Name())
	fileExt := filepath.Ext(l.FileLog.Name())

	t := time.Now()

	nowFormmated := fmt.Sprintf("%d%d%d_%d%d%d", t.Year(), t.Month(), t.Day(),
		t.Hour(), t.Minute(), t.Second())

	l.lock.Lock()
	defer l.lock.Unlock()

	if err := l.FileLog.Close(); err != nil {

		fmt.Print("error closing log file")
		return false
	}

	err := os.Rename(l.FileLog.Name(), filePath+"/"+filename+"_"+nowFormmated+fileExt)
	if err != nil {

		fmt.Println("File not able to be rotated", err)
		return false
	}

	l.FileLog, errCreatingFile = os.Create(l.FileLog.Name())

	if errCreatingFile != nil {

		fmt.Println("Error rotating the file", errCreatingFile)
		return false
	}

	l.Trace.SetOutput(l.FileLog)
	l.Debug.SetOutput(l.FileLog)
	l.Info.SetOutput(l.FileLog)
	l.Warning.SetOutput(l.FileLog)
	l.Error.SetOutput(l.FileLog)

	return true
}

//RotateLogByWeight ...
func (l *Logger) RotateLogByWeight() bool {

	fileStat, err := l.FileLog.Stat()

	if err != nil {

		fmt.Println("Error AutoRotating logs", err)
		return false
	}

	fileSize := fileStat.Size() / 1024

	if l.Rotation && uint(fileSize) >= l.RotationWeight {

		var errCreatingFile error
		filename := filepath.Base(l.FileLog.Name())
		filePath := filepath.Dir(l.FileLog.Name())
		fileExt := filepath.Ext(l.FileLog.Name())

		t := time.Now()

		nowFormmated := fmt.Sprintf("%d%d%d_%d%d%d", t.Year(), t.Month(), t.Day(),
			t.Hour(), t.Minute(), t.Second())

		l.lock.Lock()
		defer l.lock.Unlock()

		if err := l.FileLog.Close(); err != nil {

			fmt.Print("error closing log file")
			return false
		}

		err := os.Rename(l.FileLog.Name(), filePath+"/"+filename+"_"+nowFormmated+fileExt)
		if err != nil {

			fmt.Println("File not able to be rotated", err)
			return false
		}

		l.FileLog, errCreatingFile = os.Create(l.FileLog.Name())

		if errCreatingFile != nil {

			fmt.Println("Error rotating the file", errCreatingFile)
			return false
		}

		l.Trace.SetOutput(l.FileLog)
		l.Debug.SetOutput(l.FileLog)
		l.Info.SetOutput(l.FileLog)
		l.Warning.SetOutput(l.FileLog)
		l.Error.SetOutput(l.FileLog)

		return true
	}

	return true
}

func (l *Logger) intiCustomLoggers(flags int) {

	if flags&TRACE != 0 {

		l.Trace = log.New(l.FileLog, "TRACE: ", log.Ldate|log.Ltime|log.Lshortfile)
	} else {

		l.Trace = log.New(ioutil.Discard, "TRACE: ", log.Ldate|log.Ltime|log.Lshortfile)
	}

	if flags&DEBUG != 0 {

		l.Debug = log.New(l.FileLog, "DEBUG: ", log.Ldate|log.Ltime|log.Lshortfile)

	} else {

		l.Debug = log.New(ioutil.Discard, "DEBUG: ", log.Ldate|log.Ltime|log.Lshortfile)
	}

	if flags&INFO != 0 {

		l.Info = log.New(l.FileLog, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)

	} else {

		l.Info = log.New(ioutil.Discard, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)

	}

	if flags&WARN != 0 {

		l.Warning = log.New(l.FileLog, "WARN: ", log.Ldate|log.Ltime|log.Lshortfile)

	} else {

		l.Warning = log.New(ioutil.Discard, "WARN: ", log.Ldate|log.Ltime|log.Lshortfile)

	}

	if flags&ERROR != 0 {

		l.Error = log.New(l.FileLog, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
	} else {

		l.Error = log.New(ioutil.Discard, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)

	}

}

func (l *Logger) initProductionLoggers() {

	l.Trace = log.New(ioutil.Discard, "TRACE: ", log.Ldate|log.Ltime|log.Lshortfile)

	l.Debug = log.New(ioutil.Discard, "DEBUG: ", log.Ldate|log.Ltime|log.Lshortfile)

	l.Info = log.New(ioutil.Discard, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)

	l.Warning = log.New(l.FileLog, "WARN: ", log.Ldate|log.Ltime|log.Lshortfile)

	l.Error = log.New(l.FileLog, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)

}

func (l *Logger) initDevelopmentLoggers() {

	l.Trace = log.New(ioutil.Discard, "TRACE: ", log.Ldate|log.Ltime|log.Lshortfile)

	l.Debug = log.New(l.FileLog, "DEBUG: ", log.Ldate|log.Ltime|log.Lshortfile)

	l.Info = log.New(l.FileLog, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)

	l.Warning = log.New(l.FileLog, "WARN: ", log.Ldate|log.Ltime|log.Lshortfile)

	l.Error = log.New(l.FileLog, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)

}

func handleLogFile(filePath string) *os.File {

	file, err := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {

		fmt.Println("Failed to open/create log file", filePath)
		return nil
	}

	return file

}
