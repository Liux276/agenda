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
	"strings"

	"github.com/spf13/cobra"
	"github.com/sysu-615/agenda/entity"
	"github.com/sysu-615/agenda/models"
)

//退出的会议名称
var quitMeetingTitle string

// mtquitCmd represents the mtquit command
var mtquitCmd = &cobra.Command{
	Use:   "mtquit",
	Short: "Use mtquit to quit a meeting which you are the participator.",
	Long:  `Use mtquit to quit a meeting which you are the participator. [mtquit -t meeting's name]`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("mtquit called")
		login, loginUser := entity.IsLoggedIn()
		// 是否已经登录
		if !login {
			fmt.Println("Please sign in before quit a meeting")
			os.Exit(0)
		}
		models.Logger.SetPrefix("[agenda mtquit]")
		//输入会议名称为空
		if quitMeetingTitle == "" {
			models.Logger.Println("Quit meeting", quitMeetingTitle, "failed!")
			fmt.Println("The meeting's title cann't be empty!")
			os.Exit(0)
		}
		// 查找对应会议
		meetings := entity.ReadMeetingFromFile()
		for i, meeting := range meetings {
			//找到该会议
			if meeting.Title == quitMeetingTitle {
				//当前用户为会议创始人
				if meeting.Originator == loginUser.Username {
					models.Logger.Println("Quit meeting", quitMeetingTitle, "failed!")
					fmt.Println("You cann't quit the meeting you created! Please use the command mtcancel to cancel the meeting you created.")
					os.Exit(0)
				} else {
					participators := strings.Split(meeting.Participants, ",")
					for j, participator := range participators {
						//是会议参与者
						if participator == loginUser.Username {
							//退出会议
							if loginUser.Username == meeting.Participants {
								//唯一的参与者，删除会议
								newMeetingRecord := meetings[:i]
								for k := i + 1; k < len(meetings); k++ {
									newMeetingRecord = append(newMeetingRecord, meetings[k])
								}
								entity.WriteMeetingToFile(newMeetingRecord)
							} else {
								//从参与者中删除当前用户
								newParticipators := participators[:j]
								for l := j + 1; l < len(participators); l++ {
									newParticipators = append(newParticipators, participators[l])
								}
								meetings[i].Participants = strings.Join(newParticipators, ",")
								entity.WriteMeetingToFile(meetings)
							}
							models.Logger.Println("Quit meeting", quitMeetingTitle, "success!")
							fmt.Println("Success to quit the meeting", quitMeetingTitle)
							os.Exit(0)
						}
					}
					//当前用户不是该会议参与者
					models.Logger.Println("Quit meeting", quitMeetingTitle, "failed!")
					fmt.Println("Failed to quit the meeting", quitMeetingTitle, "since you are not the participator!")
					os.Exit(0)
				}
			}
		}
		//未找到会议
		models.Logger.Println("Quit meeting", quitMeetingTitle, "failed!")
		fmt.Println("Failed to quit the meeting", quitMeetingTitle, "which is not exit.")
		os.Exit(0)
	},
}

func init() {
	rootCmd.AddCommand(mtquitCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// mtquitCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// mtquitCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	mtquitCmd.Flags().StringVarP(&quitMeetingTitle, "The meeting's title you want to quit", "t", "", "The meeting's title you want to quit, you cann't be the originator.")
}
