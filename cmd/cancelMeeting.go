// Copyright © 2017 NAME HERE <EMAIL ADDRESS>
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
	"github.com/James-Yip/Agenda/service"

	"github.com/spf13/cobra"
)

// cancelMeetingCmd represents the cancelMeeting command
var cancelMeetingCmd = &cobra.Command{
	Use:   "cancelMeeting",
	Short: "Cancel a meeting",
	Long:  `Cancel a meeting created by current login user.`,
	Run: func(cmd *cobra.Command, args []string) {
		title, _ := cmd.Flags().GetString("title")
		// fmt.Println("title: " + title)
		service.CancelMeeting(title)
	},
}

func init() {
	RootCmd.AddCommand(cancelMeetingCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// cancelMeetingCmd.PersistentFlags().String("foo", "", "A help for foo")

	// local flags for quitMeetingCmd
	cancelMeetingCmd.Flags().StringP("title", "t", "", "Title of the meeting you intend to quit")
}
