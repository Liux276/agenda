package models

import (
	"os"
)

var UsersHandler *os.File
var MeetingsHandler *os.File
var partakesHandler *os.File

func init(){
	var err error
	UsersHandler, err = os.OpenFile("github.com/sysu-615/agenda/storage/users.json", os.O_RDONLY | os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	MeetingsHandler, _ = os.OpenFile("github.com/sysu-615/agenda/storage/meetings.json", os.O_RDWR, 0666)
	partakesHandler, _ = os.OpenFile("github.com/sysu-615/agenda/storage/partakes.json", os.O_RDWR, 0666)
}