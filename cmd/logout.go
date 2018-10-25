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

// logoutCmd represents the logout command
var logoutCmd = &cobra.Command{
	Use:   "logout",
	Short: "This command can logout user",
	Long:  `You can use agenda logout to logout user`,
	Run: func(cmd *cobra.Command, args []string) {
		models.Logger.SetPrefix("[agenda logout]")

		isLoggedIn, user := entity.IsLoggedIn()
		if isLoggedIn == true {
			entity.ClearCurUserInfo()
			fmt.Println(user.Username, "log out")
			models.Logger.Println(user.Username, "log out")
		} else {
			fmt.Println("No user login")
		}
		// users := entity.ReadUserInfoFromFile()

		// for i, user := range users {
		// 	if user.Login {
		// 		fmt.Println(user.Username, "log out.")
		// 		models.Logger.Println(user.Username, "log out.")
		// 		users[i].Login = false
		// 		entity.WriteUserInfoToFile(users)
		// 		return
		// 	}
		// }
		// fmt.Println("No user login")
	},
}

func init() {
	rootCmd.AddCommand(logoutCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// logoutCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// logoutCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
