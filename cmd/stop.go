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

// stopCmd represents the stop command
var stopCmd = &cobra.Command{
	Use:   "stop",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		err := vmcontrol.StopInstance(
			os.Stdin,
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
	rootCmd.AddCommand(stopCmd)
	initConfig()

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// stopCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// stopCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
