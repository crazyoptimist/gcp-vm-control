/*
Copyright © 2024 crazyoptimist <hey@crazyoptimist.net>
*/
package cmd

import (
	"log"
	"os"

	"gcp-vm-control/pkg/vmcontrol"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// startCmd represents the start command
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Starts the VM",
	Long:  `Starts the VM`,
	Run: func(cmd *cobra.Command, args []string) {
		initConfig()
		err := vmcontrol.StartInstance(
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
	rootCmd.AddCommand(startCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// startCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// startCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
