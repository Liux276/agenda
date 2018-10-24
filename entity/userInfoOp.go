package entity

import (
	"bufio"
	"fmt"
	"github.com/json-iterator/go"
	"github.com/sysu-615/agenda/models"
	"io"
	"log"
	"os"
)

type User models.User

func ReadUserInfoFromFile() []User {
	var list []User
	file, err := os.OpenFile("agenda/storage/users.json", os.O_RDWR|os.O_CREATE, 0644)
	defer file.Close()
	if err != nil {
		panic(err)
	}
	var user User
	reader := bufio.NewReader(file)
	for {
		data, errR := reader.ReadBytes('\n')
		err = jsoniter.Unmarshal(data, &user)
		if errR != nil {
			if errR == io.EOF {
				break
			} else {
				os.Stderr.Write([]byte("Read bytes from reader fail\n"))
				os.Exit(0)
			}
		}
		fmt.Println(user)
		list = append(list, user)
	}
	return list
}

func WriteUserInfoToFile(list []User) {
	file, err := os.OpenFile("../storage/users.json", os.O_RDWR|os.O_CREATE, 0644)
	defer file.Close()
	if err != nil {
		panic(err)
	}
	writer := bufio.NewWriter(file)
	var jsoniter = jsoniter.ConfigCompatibleWithStandardLibrary
	for _, user := range list {
		// 序列化
		data, err := jsoniter.Marshal(&user)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(data))
		_, errW := writer.Write([]byte(string(data)))
		writer.WriteByte('\n')
		if errW != nil {
			fmt.Println(errW)
		}
		writer.Flush()
	}
}
