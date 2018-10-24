package entity

import (
	"bufio"
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
		if len(data) <= 2 {
			continue
		}

		if data[len(data)-2] == ',' {
			data = data[0 : len(data)-2]
		}

		err = jsoniter.Unmarshal(data, &user)
		if err != nil {
			panic(err)
		}
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
			panic(errW)
		}
		writer.Flush()
	}
	writer.WriteByte(']')
	writer.WriteByte('\n')
	writer.Flush()
}
