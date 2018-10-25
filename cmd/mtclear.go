// Copyright © 2018 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/sysu-615/agenda/entity"
	"github.com/sysu-615/agenda/models"
)

// mtclearCmd represents the mtclear command
var mtclearCmd = &cobra.Command{
	Use:   "mtclear",
	Short: "Use mtclear to Cancle all the meetings you have created",
	Long:  `Use mtclear to Cancle all the meetings you have created [agenda mtclear]`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("mtcancel called")
		login, loggedUser := entity.IsLoggedIn()
		// 是否已经登录
		if !login {
			fmt.Println("Please sign in before cancle a meeting")
			os.Exit(0)
		}
		models.Logger.SetPrefix("[agenda mtclear]")

		//清除所有该用户为发起人的会议
		meetings := entity.ReadMeetingFromFile()
		newMeetingRecord := make([]models.Meeting, 0)

		for _, meeting := range meetings {
			//如果该用户不为发起人，则不清除
			if meeting.Originator != loggedUser.Username {
				newMeetingRecord = append(newMeetingRecord, meeting)
			}
		}
		entity.WriteMeetingToFile(newMeetingRecord)
		fmt.Println("All meetings of user", loggedUser.Username, "are cancelled!")
		models.Logger.Println("Cancel meetings success of user:", loggedUser.Username)
		os.Exit(0)
	},
}

func init() {
	rootCmd.AddCommand(mtclearCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// mtclearCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// mtclearCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
