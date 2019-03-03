// Copyright © 2019 Magic Lea <lea@it-easy.tw>
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
	"strconv"
	"time"

	"github.com/spf13/cobra"
)

// timestampCmd represents the timestamp command
var timestampCmd = &cobra.Command{
	Use:     "timestamp",
	Aliases: []string{"ts"},
	Short:   "Timestamp Conversion Tools",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 1 {
			// query 1: timestamp -> human readable date
			var timeMs int64
			timestamp, err := strconv.ParseInt(args[0], 10, 64)
			if err != nil {
				er(err)
			}

			if timestamp < 1e11 {
				timeMs = timestamp * 1000
			} else {
				fmt.Println("Assuming that this timestamp is in milliseconds.")
				timeMs = timestamp
			}

			fmt.Printf("Epoch to human readable time is:\t%s\n", time.Unix(timeMs/1000, timeMs%1000))
			return
		}

		yr, mon, day, hr, min, sec := 1970, 1, 1, 0, 0, 0
		var err error
		if len(args) > 0 {
			if yr, err = strconv.Atoi(args[0]); err != nil {
				er(err)
			}
		}
		if len(args) > 1 {
			if mon, err = strconv.Atoi(args[1]); err != nil {
				er(err)
			}
		}
		if len(args) > 2 {
			if day, err = strconv.Atoi(args[2]); err != nil {
				er(err)
			}
		}
		if len(args) > 3 {
			if hr, err = strconv.Atoi(args[3]); err != nil {
				er(err)
			}
		}
		if len(args) > 4 {
			if min, err = strconv.Atoi(args[4]); err != nil {
				er(err)
			}
		}
		if len(args) > 5 {
			if sec, err = strconv.Atoi(args[5]); err != nil {
				er(err)
			}
		}
		fmt.Printf("%04d/%02d/%02d %02d:%02d:%02d %v -> %v\n",
			yr, mon, day, hr, min, sec, time.Local, time.Date(yr, time.Month(mon), day, hr, min, sec, 0, time.Local).Unix())

	},
}

func init() {
	rootCmd.AddCommand(timestampCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// timestampCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// timestampCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
