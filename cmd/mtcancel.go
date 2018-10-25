// Copyright © 2018 Title HERE <EMAIL ADDRESS>
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

var meetingTitle string

// mtcancelCmd represents the mtcancel command
var mtcancelCmd = &cobra.Command{
	Use:   "mtcancel",
	Short: "Cancel a meeting by meeting Title",
	Long:  `If you are the creater of the meeting, you can cancel the meeting by the Title of the meeting.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("mtcancel called")
		login, loggedUser := entity.IsLoggedIn()
		// 是否已经登录
		if !login {
			fmt.Println("Please sign in before cancle a meeting")
			os.Exit(0)
		}
		models.Logger.SetPrefix("[agenda mtcancel]")

		//检查会议是否存在
		meetings := entity.ReadMeetingFromFile()
		for i, meeting := range meetings {
			if meeting.Title == meetingTitle {
				//检查该用户是否为发起人
				if meeting.Originator != loggedUser.Username {
					models.Logger.Println("Failed to delete meeting: ", meetingTitle)
					fmt.Println("You are not the originator of the meeting and have no permission to cancel it!")
					os.Exit(0)
				} else {
					//将该会议删除
					newMeetingRecord := make([]models.Meeting, 0)
					if i-1 >= 0 {
						newMeetingRecord = meetings[:i-1]
					}
					for j := i + 1; j < len(meetings); j++ {
						newMeetingRecord = append(newMeetingRecord, meetings[j])
					}
					for _, x := range newMeetingRecord {
						fmt.Println(x)
					}
					entity.WriteMeetingToFile(newMeetingRecord)
					fmt.Println("the meeting", meetingTitle, "are cancelled!")
					models.Logger.Println("Cancel meeting success: ", meetingTitle)
					os.Exit(0)
				}
			}
		}

		//会议不存在
		models.Logger.Println("Failed to delete meeting:", meetingTitle)
		fmt.Println("The meeting", meetingTitle, "is not exits!")
		os.Exit(0)
	},
}

func init() {
	rootCmd.AddCommand(mtcancelCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// mtcancelCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// mtcancelCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	mtcancelCmd.Flags().StringVarP(&meetingTitle, "meetingTitle", "t", "", "The title of the meeting which you want to cancel.")
}
