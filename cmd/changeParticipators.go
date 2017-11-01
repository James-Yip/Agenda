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

// changeParticipatorsCmd represents the changeParticipators command
var changeParticipatorsCmd = &cobra.Command{
	Use:   "changeParticipators",
	Short: "Change(add/delete) meeting participators.",
	Long:  `Change(add/delete) participators of a meeting created by current login user.`,
	Run: func(cmd *cobra.Command, args []string) {
		title, _ := cmd.Flags().GetString("title")
		delete_participators_str, _ := cmd.Flags().GetString("delete")
		add_participators_str, _ := cmd.Flags().GetString("add")
		// fmt.Println("title: " + title)
		// fmt.Println("delete_participators: " + delete_participators_str)
		// fmt.Println("add_participators: " + add_participators_str)

		if add_participators_str != "" {
			service.AddParticipators(title, add_participators_str)
		} else {
			service.DeleteParticipators(title, delete_participators_str)
		}
	},
}

func init() {
	RootCmd.AddCommand(changeParticipatorsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// changeParticipatorsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// local flags for changeParticipatorsCmd
	changeParticipatorsCmd.Flags().StringP("title", "t", "", "title")
	changeParticipatorsCmd.Flags().StringP("delete", "d", "", "participators that you intend to delete")
	changeParticipatorsCmd.Flags().StringP("add", "a", "", "participators that you intend to add")
}
