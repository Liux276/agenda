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

var apTitle, apParticipators string

// apCmd represents the ap command
var apCmd = &cobra.Command{
	Use:   "ap",
	Short: "Use ap to add participators to a meeting you created",
	Long:  `Use ap to add participators to a meeting you created. [agenda ap -t MeetingTitle -p Participators]`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("ap called")
		login, loginUser := entity.IsLoggedIn()
		// 是否已经登录
		if !login {
			fmt.Println("Please sign in before add participators to a meeting")
			os.Exit(0)
		}
		models.Logger.SetPrefix("[agenda ap]")
		//参数不能为空
		if apTitle == "" {
			models.Logger.Println("Add participators", apParticipators, "to meeting", apTitle, "Failed!")
			fmt.Println("the title of the meeting you want to add participators cann't be empty!")
			os.Exit(0)
		} else if apParticipators == "" {
			models.Logger.Println("Add participators", apParticipators, "to meeting", apTitle, "Failed!")
			fmt.Println("the participators you want to add cann't be empty!")
			os.Exit(0)
		}
		//查找对应会议
		meetings := entity.ReadMeetingFromFile()
		for i, meeting := range meetings {
			if meeting.Title == apTitle {
				if meeting.Originator == loginUser.Username {
					participators := strings.Split(meeting.Participants, ",")
					addParticipators := strings.Split(apParticipators, ",")
					for _, addParticipator := range addParticipators {
						pos := -1
						for j, participator := range participators {
							//已经为参与者
							if addParticipator == participator {
								fmt.Println(addParticipator, "have participated this meeting!")
								pos = j
								break
							}
						}
						//不是参与者
						if pos == -1 {
							//检查参与者的时间是否有冲突
							participatedMeetings := entity.FetchMeetingsByName(addParticipator)
							for _, participatedMeeting := range participatedMeetings {
								if (participatedMeeting.StartTime >= meeting.StartTime && participatedMeeting.StartTime < meeting.EndTime) || (participatedMeeting.EndTime > meeting.StartTime && participatedMeeting.EndTime <= meeting.EndTime) {
									models.Logger.Println("Failed to add users:", apParticipators)
									fmt.Println("Some meetings of the participator(", addParticipator, ")conflict with the meeting in terms of time")
									os.Exit(0)
								}
							}
							participators = append(participators, addParticipator)
						}
					}
					meetings[i].Participants = strings.Join(participators, ",")
					entity.WriteMeetingToFile(meetings)
					models.Logger.Println("Add participators", apParticipators, "to meeting", apTitle, "success!")
					fmt.Println("Add participators", apParticipators, "to meeting", apTitle, "success!")
					os.Exit(0)
				}
				//不是创建者
				models.Logger.Println("Add participators", apParticipators, "to meeting", apTitle, "Failed!")
				fmt.Println("Add participators", apParticipators, "to meeting", apTitle, "Failed! You are not the Originator!")
				os.Exit(0)
			}
		}
		//未找到对应会议
		models.Logger.Println("Add participators", apParticipators, "to meeting", apTitle, "Failed!")
		fmt.Println("Add participators", apParticipators, "to meeting", apTitle, "Failed! There does not have this meeting!")
		os.Exit(0)
	},
}

func init() {
	rootCmd.AddCommand(apCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// apCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// apCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	apCmd.Flags().StringVarP(&apTitle, "meeting's title", "t", "", "the meeting's title which you created and want to add participators.")
	apCmd.Flags().StringVarP(&apParticipators, "meeting's participators", "p", "", "participators you want to add to a meeting you created.")
}
