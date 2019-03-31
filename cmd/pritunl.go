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

// NOTE: customized from https://github.com/cghdev/gotunl-vpn

package cmd

import (
	"fmt"
	"os"
	"runtime"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/magiclea/gotunl"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
	"github.com/tidwall/gjson"
)

var color = map[string]string{
	"red":   "\x1b[31;1m",
	"green": "\x1b[32;1m",
	"reset": "\x1b[0m"}

type connections [][]string

var (
	password      *string
	disconnectAll *bool
)

// pritunlCmd represents the pritunl command
var pritunlCmd = &cobra.Command{
	Use:   "pritunl",
	Short: "Pritunl Client CLI",
	Run: func(cmd *cobra.Command, args []string) {
		gt := gotunl.New()

		// list profiles
		if len(gt.Profiles) == 0 {
			fmt.Println("No profiles found in Pritunl")
			os.Exit(1)
		}
		cons := gt.GetConnections()
		c := connections{}
		stdis := ""
		stcon := ""
		for pid, p := range gt.Profiles {
			if runtime.GOOS != "windows" {
				stdis = color["red"] + "Disconnected" + color["reset"]
				stcon = color["green"] + "Connected" + color["reset"]
			} else {
				stdis = "Disconnected"
				stcon = "Connected"
			}
			status := stdis
			if strings.Contains(cons, pid) {
				status = strings.Title(gjson.Get(cons, pid+".status").String())
				if status == "Connected" {
					status = stcon
				}
			}
			ptmp := []string{strconv.Itoa(p.ID), gjson.Get(p.Conf, "name").String(), status}
			c = append(c, ptmp)
			sort.Slice(c, func(i, j int) bool { return c[i][0] < c[j][0] })
		}
		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"ID", "Name", "Status"})
		table.SetAutoFormatHeaders(false)
		for _, p := range c {
			table.Append(p)
		}
		table.Render()

		// choose profile
		var id string
		fmt.Printf("Enter Profile ID or Name: ")
		fmt.Scanln(&id)

		// disconnect all connection if needed
		if *disconnectAll {
			gt.StopConnections()
			fmt.Println("sent request to stop all connections")
			time.Sleep(300 * time.Millisecond)
		}

		// connect
		for pid, p := range gt.Profiles {
			if id == gjson.Get(p.Conf, "name").String() || id == strconv.Itoa(p.ID) {
				if *password == "" {
					gt.ConnectProfile(pid, "", "")
				} else {
					gt.ConnectProfile(pid, "pritunl", *password)
				}
				fmt.Println("sent request to connect new one")
				break
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(pritunlCmd)

	password = pritunlCmd.Flags().String("password", "", "specify password")
	disconnectAll = pritunlCmd.Flags().Bool("disconnectAll", false, "whether disconnect all connections before connect new one")
}
