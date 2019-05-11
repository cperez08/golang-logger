package gologger_test

import (
	"os"
	"testing"
	"time"

	gologger "github.com/cperez08/golang-logger"
)

func TestLoggin(t *testing.T) {

	log := gologger.NewLogger("./tmp/testfile.log", gologger.DEV, 0)

	log.Info.Println("Printing DEV INFO to test")
	log.Trace.Print("Printing DEV TRACE to test")
	log.Debug.Print("Printing DEV DEBUG to test")
	log.Warning.Print("Printing DEV WARN to test")
	log.Error.Print("Printing DEV ERR to test")

	log.Info.Println("")

	log = gologger.NewLogger("./tmp/testfile.log", gologger.PROD, 0)

	log.Info.Println("Printing PROD INFO to test")
	log.Trace.Print("Printing PROD TRACE to test")
	log.Debug.Print("Printing PROD DEBUG to test")
	log.Warning.Print("Printing PROD WARN to test")
	log.Error.Print("Printing PROD ERR to test")

	log.Error.Println("")

	log = gologger.NewLogger("./tmp/testfile.log", gologger.CUSTOM, gologger.TRACE|gologger.DEBUG)

	log.Info.Println("Printing CUSTOM INFO to test")
	log.Trace.Print("Printing CUSTOM TRACE to test")
	log.Debug.Print("Printing CUSTOM DEBUG to test")
	log.Warning.Print("Printing CUSTOM WARN to test")
	log.Error.Print("Printing CUSTOM ERR to test")
}

func TestLoggerWithRotation(t *testing.T) {

	log := gologger.NewLoggerWithRotation("./tmp/testfile.log", gologger.DEV, 0, 10)

	for i := 0; i < 40; i++ {

		log.Info.Println("Printing DEV AUTO INFO to test", i)
		log.Trace.Println("Printing DEV AUTO TRACE to test", i)
		log.Debug.Println("Printing DEV AUTO DEBUG to test", i)
		log.Warning.Println("Printing DEV AUTO WARN to test", i)
		log.Error.Println("Printing DEV AUTO ERR to test", i)
		time.Sleep(time.Millisecond * 50)
		log.RotateLogByWeight()

	}

	log = gologger.NewLoggerWithRotation("./tmp/testfile.log", gologger.PROD, 0, 1)

	log.Info.Println("Printing PROD AUTO INFO to test")
	log.Trace.Println("Printing PROD AUTO TRACE to test")
	log.Debug.Println("Printing PROD AUTO DEBUG to test")
	log.Warning.Println("Printing PROD AUTO WARN to test")
	log.Error.Println("Printing PROD AUTO ERR to test")

	log = gologger.NewLoggerWithRotation("./tmp/testfile.log", gologger.CUSTOM, gologger.TRACE|gologger.INFO, 1)

	log.Info.Println("Printing CUSTOM AUTO INFO to test")
	log.Trace.Print("Printing CUSTOM AUTO TRACE to test")
	log.Debug.Print("Printing CUSTOM AUTO DEBUG to test")
	log.Warning.Print("Printing CUSTOM AUTO WARN to test")
	log.Error.Print("Printing CUSTOM AUTO ERR to test")

}

func TestLoggerWithManualRotation(t *testing.T) {

	log := gologger.NewLoggerWithRotation("./tmp/testfile.log", gologger.DEV, 0, 10)

	log.Info.Println("Printing DEV AUTO INFO to test")
	log.Trace.Println("Printing DEV AUTO TRACE to test")
	log.Debug.Println("Printing DEV AUTO DEBUG to test")
	log.Warning.Println("Printing DEV AUTO WARN to test")
	log.Error.Println("Printing DEV AUTO ERR to test")
	time.Sleep(time.Millisecond * 50)
	log.RotateLog()
}

func TestLoggerWithOpenFile(t *testing.T) {

	file, err := os.OpenFile("./tmp/logger_opefile.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {

		t.Fail()
		return
	}

	defer file.Close()

	log := gologger.NewLoggerWithOpenFile(file, gologger.DEV, 0)

	log.Info.Println("Printing DEV INFO WITH  OPEN FILE ")
	log.Trace.Print("Printing DEV TRACE  WITH  OPEN FILE")
	log.Debug.Print("Printing DEV DEBUG  WITH  OPEN FILE")
	log.Warning.Print("Printing DEV WARN  WITH  OPEN FILE")
	log.Error.Print("Printing DEV ERR  WITH  OPEN FILE")

	log = gologger.NewLoggerWithOpenFile(file, gologger.PROD, 0)

	log.Info.Println("Printing PROD INFO WITH  OPEN FILE ")
	log.Trace.Print("Printing PROD TRACE  WITH  OPEN FILE")
	log.Debug.Print("Printing PROD DEBUG  WITH  OPEN FILE")
	log.Warning.Print("Printing PROD WARN  WITH  OPEN FILE")
	log.Error.Print("Printing PROD ERR  WITH  OPEN FILE")

	log = gologger.NewLoggerWithOpenFile(file, gologger.CUSTOM, gologger.INFO|gologger.DEBUG)

	log.Info.Println("Printing CUSTOM INFO WITH  OPEN FILE ")
	log.Trace.Print("Printing CUSTOM TRACE  WITH  OPEN FILE")
	log.Debug.Print("Printing CUSTOM DEBUG  WITH  OPEN FILE")
	log.Warning.Print("Printing CUSTOM WARN  WITH  OPEN FILE")
	log.Error.Print("Printing CUSTOM ERR  WITH  OPEN FILE")
}

func TestMakingFails(t *testing.T) {

	log := gologger.NewLoggerWithOpenFile(nil, gologger.DEV, 0)

	if log != nil {

		t.Fail()
	}

	log2 := gologger.NewLogger("", gologger.DEV, 0)

	if log2 != nil {

		t.Fail()
	}

	log3 := gologger.NewLoggerWithRotation("", gologger.DEV, 0, 1)

	if log3 != nil {

		t.Fail()
	}

	log4 := gologger.NewLogger("./tmp/testfile.log", gologger.CUSTOM, gologger.WARN|gologger.ERROR)

	if closed := log4.CloseFile(); closed == false {

		t.Fail()
	}

}
