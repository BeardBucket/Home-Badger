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
	"github.com/BeardBucket/Home-Badger/src/mainz"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var addonConfigFile string
var addonPort int

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use:   "run",
	Short: "OnRun the daemon",
	Run: func(cmd *cobra.Command, args []string) {
		// Exit channel
		onSIGTERM := make(chan os.Signal)
		signal.Notify(onSIGTERM, os.Interrupt, syscall.SIGTERM)

		// Send SIGTERM so that we exit out. Print the problem too.
		notify := func(msg string, err error) {
			onSIGTERM <- syscall.SIGTERM
		}

		// OnExit when all done - Called on interrupt
		go func() {
			<-onSIGTERM
			mainz.OnExit()
			os.Exit(1)
		}()

		// Run late inits
		err := mainz.OnLateInit(cmd, args, notify, viper.GetViper())
		if err != nil {
			notify("Problem during late inits", err)
		}

		// Main worker
		go func() {
			err := mainz.OnRun()
			if err != nil {
				notify("Problem in main runner", err)
			}
		}()

		// Sleep main thread forever
		for {
			err := mainz.OnCycle()
			if err != nil {
				notify("Problem in cycle runner", err)
			}
			time.Sleep(mainz.GetMain().Vpr().GetDuration("cycle.duration"))
		}
	},
}

func init() {
	// Set a default for a main function: Special case
	viper.SetDefault("cycle.duration", 1*time.Second)
	rootCmd.AddCommand(runCmd)
	runCmd.PersistentFlags().StringVar(&addonConfigFile, "addon-options", "/data/options.json", "hass add-on options")
	runCmd.PersistentFlags().IntVar(&addonPort, "port", 80, "Port to listen on")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// runCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// runCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
