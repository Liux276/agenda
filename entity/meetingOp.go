package entity

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"github.com/json-iterator/go"
	"github.com/sysu-615/agenda/models"
)

func ReadMeetingFromFile() []models.Meeting {
	var list []models.Meeting
	file, err := os.OpenFile("github.com/sysu-615/agenda/storage/meetings.json", os.O_RDONLY|os.O_CREATE, 0644)
	defer file.Close()
	if err != nil {
		panic(err)
	}
	var meeting models.Meeting
	reader := bufio.NewReader(file)
	for {
		data, errR := reader.ReadBytes('\n')
		if errR != nil {
			if errR == io.EOF {
				break
			} else {
				os.Stderr.Write([]byte("Read bytes from reader failed\n"))
				os.Exit(0)
			}
		}
		// 过滤'['、']'
		if len(data) <= 2 {
			continue
		}

		data = data[0 : len(data)-1]

		err = jsoniter.Unmarshal(data, &meeting)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// fmt.Println(user)
		list = append(list, meeting)
	}
	return list
}

func WriteMeetingToFile(list []models.Meeting) {
	file, err := os.OpenFile("github.com/sysu-615/agenda/storage/meetings.json", os.O_WRONLY|os.O_CREATE, 0644)
	defer file.Close()
	if err != nil {
		panic(err)
	}
	writer := bufio.NewWriter(file)

	var jsoniter = jsoniter.ConfigCompatibleWithStandardLibrary

	writer.WriteByte('[')
	writer.WriteByte('\n')
	for i, meeting := range list {
		// 序列化
		data, err := jsoniter.Marshal(&meeting)

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

func FetchMeetingsByName(name string) []models.Meeting {
	var list []models.Meeting
	file, err := os.OpenFile("github.com/sysu-615/agenda/storage/meetings.json", os.O_RDONLY|os.O_CREATE, 0644)
	defer file.Close()
	if err != nil {
		panic(err)
	}
	var meeting models.Meeting
	reader := bufio.NewReader(file)
	for {
		data, errR := reader.ReadBytes('\n')
		if errR != nil {
			if errR == io.EOF {
				break
			} else {
				os.Stderr.Write([]byte("Read bytes from reader failed\n"))
				os.Exit(0)
			}
		}
		// 过滤'['、']'
		if len(data) <= 2 {
			continue
		}
		data = data[0 : len(data)-1]
		err = jsoniter.Unmarshal(data, &meeting)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		// 获取name所参加的所有会议
		if meeting.Originator == name {
			list = append(list, meeting)
		} else {
			for _, nameStr := range strings.Split(meeting.Participants, ",") {
				if nameStr == name {
					list = append(list, meeting)
					break
				}
			}
		}
	}
	return list
}
