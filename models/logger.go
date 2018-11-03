package models

import (
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

var Logger *log.Logger
var ExecPath string

func init() {
	for _, arg := range os.Args {
		if strings.Split(arg, "=")[0] == "go_test" {
			ExecPath = strings.Replace(strings.Split(arg, "=")[1], "\\", "/", -1)
			if ExecPath[len(ExecPath)-1] != '/' {
				ExecPath += "/"
			}
			return
		}
	}
	file, _ := exec.LookPath(os.Args[0])
	path, _ := filepath.Abs(file)
	index := strings.LastIndex(path, string(os.PathSeparator))
	// the current path will be $GOPATH/bin
	// so I here truncate the bin and add src to get $GOPATH/src
	ExecPath = path[:index-3] + "src/"
	ExecPath = strings.Replace(ExecPath, "\\", "/", -1)
}

func init() {
	var logHandler *os.File
	logHandler, err := os.OpenFile(ExecPath+"github.com/sysu-615/agenda/log/logFile.txt", os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	Logger = log.New(logHandler, "", log.LstdFlags)
}
