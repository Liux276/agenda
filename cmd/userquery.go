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

	"github.com/spf13/cobra"
	"github.com/sysu-615/agenda/entity"
	"github.com/sysu-615/agenda/models"
)

// userqueryCmd represents the userquery command
var userqueryCmd = &cobra.Command{
	Use:   "userquery",
	Short: "This command can query all user information only for logged in users",
	Long:  `You can use agenda userquery to get all user information only for logged in users`,
	Run: func(cmd *cobra.Command, args []string) {
		models.Logger.SetPrefix("[agenda userquery]")
		isLoggedIn, user := entity.IsLoggedIn()
		if isLoggedIn == true {
			models.Logger.Println("UserQuery", user.Username, "query all users infomation!")
			users := entity.ReadUserInfoFromFile()
			fmt.Println("Name\tPhone\t\tEmail")
			for _, userInfo := range users {
				fmt.Printf("%-8s%-16s%s\n", userInfo.Username, userInfo.Telephone, userInfo.Email)
			}
		} else {
			fmt.Println("Please login")
		}
	},
}

func init() {
	rootCmd.AddCommand(userqueryCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// userqueryCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// userqueryCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}
