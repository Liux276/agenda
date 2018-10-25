package entity

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/json-iterator/go"
	"github.com/sysu-615/agenda/models"
)

func ReadUserInfoFromFile() []models.User {
	var list []models.User
	file, err := os.OpenFile("github.com/sysu-615/agenda/storage/users.json", os.O_RDWR|os.O_CREATE, 0644)
	defer file.Close()

	if err != nil {
		panic(err)
	}
	var user models.User
	reader := bufio.NewReader(file)
	for {
		data, errR := reader.ReadBytes('\n')
		if errR != nil {
			if errR == io.EOF {
				break
			} else {
				os.Stderr.Write([]byte("Read bytes from reader fail\n"))
				os.Exit(0)
			}
		}

		// 过滤'['、']'
		if data[0] == ']' || data[0] == '[' {
			continue
		}

		if data[len(data)-2] == ',' {
			data = data[0 : len(data)-2]
		}

		err = jsoniter.Unmarshal(data, &user)
		if err != nil {
			// fmt.Println(err)
			os.Exit(1)
		}

		// fmt.Println(user)
		list = append(list, user)
	}
	return list
}

func WriteUserInfoToFile(list []models.User) {
	file, err := os.OpenFile("github.com/sysu-615/agenda/storage/users.json", os.O_RDWR|os.O_CREATE, 0644)
	defer file.Close()

	if err != nil {
		panic(err)
	}

	writer := bufio.NewWriter(file)
	var jsoniter = jsoniter.ConfigCompatibleWithStandardLibrary

	writer.WriteByte('[')
	writer.WriteByte('\n')
	for i, user := range list {
		// 序列化
		data, err := jsoniter.Marshal(&user)
		if err != nil {
			log.Fatal(err)
		}
		writer.WriteByte('\t')
		_, errW := writer.Write([]byte(string(data)))
		if i != len(list)-1 {
			writer.WriteByte(',')
		}
		writer.WriteByte('\n')
		if errW != nil {
			fmt.Println(errW)
		}
		writer.Flush()
	}
	writer.WriteByte(']')
	writer.WriteByte('\n')
	writer.Flush()
}

func SaveCurUserInfo(loginUser models.User) {
	file, err := os.OpenFile("github.com/sysu-615/agenda/storage/curUser.txt", os.O_RDWR|os.O_CREATE, 0644)
	defer file.Close()

	if err != nil {
		panic(err)
	}

	writer := bufio.NewWriter(file)
	// 序列化
	data, err := jsoniter.Marshal(&loginUser)
	if err != nil {
		log.Fatal(err)
	}
	_, errW := writer.Write([]byte(string(data)))
	writer.WriteByte('\n')
	if errW != nil {
		panic(errW)
	}
	writer.Flush()
}

func ClearCurUserInfo() {
	err := os.Truncate("github.com/sysu-615/agenda/storage/curUser.txt", 0)
	if err != nil {
		panic(err)
	}
}

func IsLoggedIn() (bool, models.User) {
	file, err := os.OpenFile("github.com/sysu-615/agenda/storage/curUser.txt", os.O_RDWR|os.O_CREATE, 0644)
	defer file.Close()

	if err != nil {
		panic(err)
	}
	var user models.User
	// writer := bufio.NewWriter(file)
	// var jsoniter = jsoniter.ConfigCompatibleWithStandardLibrary
	reader := bufio.NewReader(file)
	data, errR := reader.ReadBytes('\n')
	if errR != nil {
		if errR == io.EOF {
			// 没有用户登陆
			return false, user
		} else {
			os.Stderr.Write([]byte("Read bytes from reader fail\n"))
			os.Exit(0)
		}
	} else {
		// 已经登陆
		err = jsoniter.Unmarshal(data, &user)
		if err != nil {
			panic(err)
		}
	}
	return true, user
}

func IsUser(name string) bool {
	users := ReadUserInfoFromFile()
	for _, user := range users {
		if user.Username == name {
			return true
		}
	}
	return false
}
