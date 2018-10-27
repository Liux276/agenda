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

var queryStartTime, queryEndTime string

// mtqueryCmd represents the mtquery command
var mtqueryCmd = &cobra.Command{
	Use:   "mtquery",
	Short: "Use mtquery to query the meetings in the Time Range start to end",
	Long:  `Use mtquery to query the meetings in the Time Range start to end [agenda mtquery -s startTime -e endTime ]`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("mtquery called")
		login, loginUser := entity.IsLoggedIn()
		// 是否已经登录
		if !login {
			fmt.Println("Please sign in before query a meeting")
			os.Exit(0)
		}
		models.Logger.SetPrefix("[agenda mtquery]")
		//判断参数是否合法
		if queryEndTime == "" || queryStartTime == "" {
			fmt.Println("The start time and end time cann't be empty!")
			os.Exit(0)
		} else if queryEndTime < queryStartTime {
			fmt.Println("The start time cann't be greater than the end time!")
			os.Exit(0)
		}
		//查找会议
		meetings := entity.FetchMeetingsByName(loginUser.Username)
		var queryedMeetings []models.Meeting
		for _, meeting := range meetings {
			if (meeting.StartTime >= queryStartTime && meeting.StartTime < queryEndTime) || (meeting.EndTime > queryStartTime && meeting.EndTime <= queryEndTime) {
				queryedMeetings = append(queryedMeetings, meeting)
			}
		}
		if len(queryedMeetings) == 0 {
			//查找失败
			models.Logger.Println("Failed to find the meeting between:", queryStartTime, queryEndTime)
			fmt.Println("Failed to find the meetings. You have no meetings between", queryStartTime, queryEndTime)
		} else {
			models.Logger.Println("Success to find the meeting between:", queryStartTime, queryEndTime)
			fmt.Println("Success to find the meeting between:", queryStartTime, queryEndTime)
			for i, meeting := range queryedMeetings {
				fmt.Println("[", i+1, "]:", meeting.Title, meeting.Originator, meeting.StartTime, meeting.EndTime, meeting.Participants)
			}
			os.Exit(0)
		}
	},
}

func init() {
	rootCmd.AddCommand(mtqueryCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// mtqueryCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// mtqueryCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	mtqueryCmd.Flags().StringVarP(&queryStartTime, "start time of meetings", "s", "", "the start time of meetings you want to query")
	mtqueryCmd.Flags().StringVarP(&queryEndTime, "end time of meetings", "e", "", "the end time of meetings you want to query")
}
