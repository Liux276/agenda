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

var canceledMeetingTitle string

// mtcancelCmd represents the mtcancel command
var mtcancelCmd = &cobra.Command{
	Use:   "mtcancel",
	Short: "Use mtcancel to cancel a meeting by meeting's Title",
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
		//检查输入的会议名称
		if canceledMeetingTitle == "" {
			fmt.Println("The meeting's title cann't be empty! Please enter the title of the meeting you want to cancel.")
			os.Exit(0)
		}
		//检查会议是否存在
		meetings := entity.ReadMeetingFromFile()
		for i, meeting := range meetings {
			if meeting.Title == canceledMeetingTitle {
				//检查该用户是否为发起人
				if meeting.Originator != loggedUser.Username {
					models.Logger.Println("Failed to delete meeting: ", canceledMeetingTitle)
					fmt.Println("You are not the originator of the meeting and have no permission to cancel it!")
					os.Exit(0)
				} else {
					//将该会议删除
					meetings = append(meetings[:i], meetings[i+1:]...)
					entity.WriteMeetingToFile(meetings)
					fmt.Println("the meeting", canceledMeetingTitle, "are cancelled!")
					models.Logger.Println("Cancel meeting success: ", canceledMeetingTitle)
					os.Exit(0)
				}
			}
		}

		//会议不存在
		models.Logger.Println("Failed to delete meeting:", canceledMeetingTitle)
		fmt.Println("The meeting", canceledMeetingTitle, "is not exits!")
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
	mtcancelCmd.Flags().StringVarP(&canceledMeetingTitle, "canceledMeetingTitle", "t", "", "The title of the meeting which you want to cancel.")
}
