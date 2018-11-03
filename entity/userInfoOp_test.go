package entity

import (
	"bufio"
	"io"
	"os"
	"testing"

	jsoniter "github.com/json-iterator/go"
	"github.com/sysu-615/agenda/models"
)

var users = []models.User{
	{
		Username:  "liuyh73",
		Password:  "123",
		Telephone: "12364579823",
		Email:     "a@163.com",
	},
	{
		Username:  "liuyh74",
		Password:  "123",
		Telephone: "12364579823",
		Email:     "a@163.com",
	},
}

func TestReadUserInfoFromFile_WriteUserInfoToFile(t *testing.T) {
	WriteUserInfoToFile(users)

	usersRead := ReadUserInfoFromFile()
	if len(usersRead) == 2 && users[0] == usersRead[0] && users[1] == usersRead[1] {
		t.Log("ReadUserInfoFromFile 和 WriteUserInfoToFile 测试通过")
	} else {
		t.Error("ReadUserInfoFromFile 或者 WriteUserInfoToFile 测试失败")
	}
}

func TestSaveCurUserInfo(t *testing.T) {
	// 假设登陆用户为liuyh73
	SaveCurUserInfo(users[0])
	file, err := os.OpenFile(models.ExecPath+"github.com/sysu-615/agenda/storage/curUser.txt", os.O_RDWR|os.O_CREATE, 0644)
	defer file.Close()

	if err != nil {
		panic(err)
	}
	var loginUser models.User
	reader := bufio.NewReader(file)
	data, _ := reader.ReadBytes('\n')

	jsoniter.Unmarshal(data, &loginUser)
	if loginUser == users[0] {
		t.Log("SaveCurUserInfo 测试通过")
	} else {
		t.Error("SaveCurUserInfo 测试失败")
	}
}

func TestIsLoggedIn(t *testing.T) {
	login, loginUser := IsLoggedIn()
	if login && loginUser == users[0] {
		t.Log("IsLoggedIn 测试通过")
	} else {
		t.Error("IsLoggedIn 测试失败")
	}
}

func TestClearCurUserInfo(t *testing.T) {
	ClearCurUserInfo()
	file, err := os.OpenFile(models.ExecPath+"github.com/sysu-615/agenda/storage/curUser.txt", os.O_RDWR|os.O_CREATE, 0644)
	defer file.Close()

	if err != nil {
		panic(err)
	}
	reader := bufio.NewReader(file)
	data, err := reader.ReadBytes('\n')
	if string(data) == "" && err == io.EOF {
		t.Log("ClearCurUserInfo 测试通过")
	} else {
		t.Error("ClearCurUserInfo 测试失败")
	}
}

func TestIsUser(t *testing.T) {
	isUser := IsUser("liuyh73")
	isNotUser := IsUser("liuyh75")
	if isUser && !isNotUser {
		t.Log("IsUser 测试通过")
	} else {
		t.Error("IsUser 测试失败")
	}
}

func TestRemoveUser(t *testing.T) {
	RemoveUser("liuyh73")
	users := ReadUserInfoFromFile()
	RemoveUser("liuyh74")
	users2 := ReadUserInfoFromFile()
	if len(users) == 1 && len(users2) == 0 {
		t.Log("RemoveUser 测试通过")
	} else {
		t.Error("RemoveUser 测试失败")
	}
}
