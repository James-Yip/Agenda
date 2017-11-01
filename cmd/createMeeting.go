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

// createMeetingCmd represents the createMeeting command
var createMeetingCmd = &cobra.Command{
	Use:   "createMeeting",
	Short: "Create a meeting",
	Long: `Create a meeting.
At least one participator should be provided.`,
	Run: func(cmd *cobra.Command, args []string) {
		title, _ := cmd.Flags().GetString("title")
		participators_str, _ := cmd.Flags().GetString("participators")
		startTime, _ := cmd.Flags().GetString("start")
		endTime, _ := cmd.Flags().GetString("end")
		// fmt.Println("title: " + title)
		// fmt.Println("participators: " + participators_str)
		// fmt.Println("startTime: " + startTime)
		// fmt.Println("endTime: " + endTime)

		service.CreateMeeting(title, participators_str, startTime, endTime)
	},
}

func init() {
	RootCmd.AddCommand(createMeetingCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createMeetingCmd.PersistentFlags().String("foo", "", "A help for foo")

	// local flags for createMeetingCmd
	createMeetingCmd.Flags().StringP("title", "t", "", "title")
	createMeetingCmd.Flags().StringP("participators", "p", "", "participators")
	createMeetingCmd.Flags().StringP("start", "s", "", "start time (format: yyyy-mm-dd/hh:mm)")
	createMeetingCmd.Flags().StringP("end", "e", "", "end time (format: yyyy-mm-dd/hh:mm)")
}
