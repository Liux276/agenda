package models

import (
	"os"
	"log"
)

var Logger *log.Logger

func init(){
	var logHandler *os.File
	logHandler, err := os.OpenFile("github.com/sysu-615/agenda/log/logFile.txt", os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	Logger = log.New(logHandler, "", log.LstdFlags)
}