## GOLANG LOGGER

This is a simple logging library for Golang that handles the output log in files, you can rotate the logs and apply different levels of logging.


## Instalation

```bash

go get github.com/cperez08/golang-logger

```

## Basic Usage

```go

import gologger "github.com/cperez08/golang-logger"

// Basic Usage with Dev Configuration

log := gologger.NewLogger("./tmp/testfile.log", gologger.DEV, 0)

log.Info.Println("Print Hello World for develop")

// Basic Usage with Prod Configuration

log := gologger.NewLogger("./tmp/testfile.log", gologger.PROD, 0)

log.Error.Println("Print Hello World Prod")

// Basic Usage with custom Configuration

log := gologger.NewLogger("./tmp/testfile.log", gologger.CUSTOM, gologger.TRACE|gologger.DEBUG|gologger.WARN|gologger.ERROR)

log.Trace.Print("Print Hello World Custom")

```

## Usage With Log Rotation

```go

import gologger "github.com/cperez08/golang-logger"

log := gologger.NewLoggerWithRotation("./tmp/testfile.log", gologger.CUSTOM, gologger.TRACE|gologger.INFO, 100)

log.Info.Println("Printing log with rotation at 100KB")

log.RotateLogByWeight()


// Also can be rotated without waiting for the weight
log.RotateLog()
```

## Usage With an Open File

```go

import gologger "github.com/cperez08/golang-logger"

file, err := os.OpenFile("./tmp/logger_opefile.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {

		t.Fail()
		return
	}

defer file.Close()

log = gologger.NewLoggerWithOpenFile(file, gologger.PROD, 0)

log.Error.Println("Print log  with an  open file")


```


##  TODO

- Create automatic rotation logs
- Improve tests
- Create MultiWriting log std.Out and Files at the same time


## Lincense

- MIT LICENSE, see LICENSE file

## Inspired by these Posts

- https://www.ardanlabs.com/blog/2013/11/using-log-package-in-go.html 
- https://stackoverflow.com/questions/28796021/how-can-i-log-in-golang-to-a-file-with-log-rotation
