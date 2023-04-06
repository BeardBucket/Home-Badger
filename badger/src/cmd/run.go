package cmd

/*
Copyright © 2023 Paulson McIntyre

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program. If not, see <http://www.gnu.org/licenses/>.
*/

import (
	"fmt"
	"github.com/BeardBucket/Home-Badger/src/mainz"
	"github.com/spf13/cobra"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Start the daemon",
	Run: func(cmd *cobra.Command, args []string) {
		c := make(chan os.Signal)
		signal.Notify(c, os.Interrupt, syscall.SIGTERM)

		notify := func(msg string, err error) {
			signal.Notify(c, os.Interrupt)
		}

		go func() {
			<-c
			mainz.Cleanup()
			os.Exit(1)
		}()

		err := mainz.LateInits()
		if err != nil {
			notify("Problem during late inits", err)
		}

		go func() {
			err := mainz.Start()
			if err != nil {
				notify("Problem in main runner", err)
			}
		}()

		for {
			fmt.Println("sleeping...")
			time.Sleep(30 * time.Second)
		}
	},
}

func init() {
	rootCmd.AddCommand(runCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// runCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// runCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
