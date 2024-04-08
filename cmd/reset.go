/*
Copyright © 2024 crazyoptimist <hey@crazyoptimist.net>
*/
package cmd

import (
	"gcp-vm-control/pkg/vmcontrol"
	"log"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// resetCmd represents the reset command
var resetCmd = &cobra.Command{
	Use:   "reset",
	Short: "Resets the VM",
	Long:  `Resets the VM`,
	Run: func(cmd *cobra.Command, args []string) {
		initConfig()
		err := vmcontrol.ResetInstance(
			os.Stdout,
			viper.GetString("ProjectID"),
			viper.GetString("Zone"),
			viper.GetString("InstanceName"),
		)
		if err != nil {
			log.Fatalln("Operation failed: ", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(resetCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// resetCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// resetCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
