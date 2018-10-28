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

var rpTitle, rpParticipators string

// rpCmd represents the rp command
var rpCmd = &cobra.Command{
	Use:   "rp",
	Short: "use rp to remove user from the meeting you created",
	Long:  `use rp to remove user from the meeting you created [agenda rp -t MeetingTitle -p Participators`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("rp called")
		login, loginUser := entity.IsLoggedIn()
		// 是否已经登录
		if !login {
			fmt.Println("Please sign in before remove participators to a meeting")
			os.Exit(0)
		}
		models.Logger.SetPrefix("[agenda rp]")
		//参数不能为空
		if rpTitle == "" {
			models.Logger.Println("Remove participators", rpParticipators, "to meeting", rpTitle, "Failed!")
			fmt.Println("the title of the meeting you want to remove participators cann't be empty!")
			os.Exit(0)
		} else if rpParticipators == "" {
			models.Logger.Println("Remove participators", rpParticipators, "to meeting", rpTitle, "Failed!")
			fmt.Println("the participators you want to remove cann't be empty!")
			os.Exit(0)
		}
		//查找对应会议
		meetings := entity.ReadMeetingFromFile()
		for i, meeting := range meetings {
			if meeting.Title == rpTitle {
				if meeting.Originator == loginUser.Username {
					participators := strings.Split(meeting.Participants, ",")
					removeParticipators := strings.Split(rpParticipators, ",")
					for _, removeParticipator := range removeParticipators {
						pos := -1
						for j, participator := range participators {
							//为参与者
							if removeParticipator == participator {
								participators = append(participators[:j], participators[j+1:]...)
								pos = j
								break
							}
						}
						//不是参与者
						if pos == -1 {
							models.Logger.Println("Failed to remve users:", rpParticipators)
							fmt.Println(removeParticipator, "is not a participator of this meeting", rpTitle)
							os.Exit(0)
						}
					}
					newParticipators := strings.Join(participators, ",")
					if newParticipators == "" {
						meetings = append(meetings[:i], meetings[i+1:]...)
					} else {
						meetings[i].Participants = newParticipators
					}
					entity.WriteMeetingToFile(meetings)
					models.Logger.Println("Remove participators", rpParticipators, "from meeting", rpTitle, "success!")
					fmt.Println("Remove participators", rpParticipators, "from meeting", rpTitle, "success!")
					os.Exit(0)
				}
				//不是创建者
				models.Logger.Println("Remove participators", rpParticipators, "to meeting", rpTitle, "Failed!")
				fmt.Println("Remove participators", rpParticipators, "to meeting", rpTitle, "Failed! You are not the Originator!")
				os.Exit(0)
			}
		}
		//未找到对应会议
		models.Logger.Println("Remove participators", rpParticipators, "to meeting", rpTitle, "Failed!")
		fmt.Println("Remove participators", rpParticipators, "to meeting", rpTitle, "Failed! There does not have this meeting!")
		os.Exit(0)
	},
}

func init() {
	rootCmd.AddCommand(rpCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// rpCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// rpCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	rpCmd.Flags().StringVarP(&rpTitle, "Meeting's title", "t", "", "The meeting's title you want to remove participators")
	rpCmd.Flags().StringVarP(&rpParticipators, "Participators", "p", "", "Participators you want to add to the meeting you created")
}
