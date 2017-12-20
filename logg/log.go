package logg

import (
	"log"
	"os"
)

type consoleFileWriter struct {
	file	*os.File
}

// 同时向文件和stdout写日志
func (con *consoleFileWriter) Write(buf []byte) (n int, err error) {
	os.Stdout.Write(buf)
	return con.file.Write(buf)
}

var Logger *log.Logger
var logFile *os.File

func init() {
	os.Mkdir("logs", 0777)
	logFile, err := os.Create("logs/monitor.log")
	if nil != err {
		panic(err)
	}

	cfWriter := &consoleFileWriter{
		file: logFile,
	}

	Logger = log.New(cfWriter, "[INFO] ", log.LstdFlags)
}



