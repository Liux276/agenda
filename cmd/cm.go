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

var createdMeeting models.Meeting

// cmCmd represents the cm command
var cmCmd = &cobra.Command{
	Use:   "cm",
	Short: "This command can create a meeting",
	Long:  `You can use agenda cm to create a meeting`,
	Run: func(cmd *cobra.Command, args []string) {
		meetings := entity.ReadMeetingFromFile()
		// 查找所有的会议，查看Title是否重复
		for _, meeting := range meetings {
			if meeting.Title == createdMeeting.Title {
				models.Logger.Println("The meeting's title", meeting.Title, "has been occupied")
				fmt.Println("The meeting's title", meeting.Title, "has been occupied")
				os.Exit(0)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(cmCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// cmCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// cmCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	cmCmd.Flags().StringVarP(&createdMeeting.Title, "title", "t", "", "The Meeting's Title")
	cmCmd.Flags().StringVarP(&createdMeeting.Originator, "originator", "o", "", "The Meeting's Originator")
	cmCmd.Flags().StringVarP(&createdMeeting.Participants, "participants", "P", "", "The Meeting's Participants")
	cmCmd.Flags().StringVarP(&createdMeeting.StartTime, "startTime", "s", "", "The Meeting's StartTime")
	cmCmd.Flags().StringVarP(&createdMeeting.Endtime, "endtime", "e", "", "The Meeting's Endtime")
}
