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
	"io"
	"encoding/json"
	"bufio"
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/sysu-615/agenda/models"
)

var user models.User

// registerCmd represents the register command
var registerCmd = &cobra.Command{
	Use:   "register",
	Short: "This command can register user",
	Long:  `You can use agenda register to sign up one user`,
	Run: func(cmd *cobra.Command, args []string) {
		usersReader := bufio.NewReader(models.UsersHandler)
		userQuery := new(models.User)
		var userBytes []byte
		var readErr, err error
		for {
			userBytes, readErr = usersReader.ReadBytes(byte(','))
			fmt.Println(userBytes)
			if readErr == io.EOF {
				break
			}
			err = json.Unmarshal(userBytes, userQuery)
			if err != nil {
				panic(err)
			}
			if user.Username == userQuery.Username {
				fmt.Println("The username", user.Username, "has been registered")
				os.Exit(1)
			}
		}

		usersWriter := bufio.NewWriter(models.UsersHandler)
		userBytes, err = json.Marshal(user) 
		if err != nil {
			panic(err)
		}
		fmt.Println(usersWriter)

		length, err := usersWriter.Write(userBytes)
		fmt.Println(length)
		if err != nil {
			panic(err)
		}
		models.Logger.SetPrefix("[agenda register]")
		models.Logger.Println("Register", user.Username, "successfully!")
	},
}

func init() {
	rootCmd.AddCommand(registerCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// registerCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// registerCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	
	registerCmd.Flags().StringVarP(&user.Username, "username", "u", "", "The User's Username")
	registerCmd.Flags().StringVarP(&user.Password, "password", "p", "", "The User's Password")
	registerCmd.Flags().StringVarP(&user.Email, "email", "e", "", "The User's Email")
	registerCmd.Flags().StringVarP(&user.Telephone, "telephone", "P", "", "The User's telephone")
}
