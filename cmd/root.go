/*
Copyright © 2024 crazyoptimist <hey@crazyoptimist.net>
*/
package cmd

import (
	"log"
	"os"

	"gcp-vm-control/pkg/utils"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "gcp-vm-control",
	Short: "Control GCP VM",
	Long: `Turn on and off Google Cloud Virtual Machines.

###################################################################################
To use this application, you need to generate ADC (application default credentials)

Install the GCP CLI:
https://cloud.google.com/sdk/docs/install

Initialize the GCP CLI:
gcloud init

Generate ADC:
gcloud auth application-default login
###################################################################################
`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

var cfgFile string

const DEFAULT_CONFIG_FILE_NAME = "gcp-vm-control"

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.config/gcp-vm-control.toml)")
}

func initConfig() {
	// Don't forget to read config either from cfgFile or from default config directory!
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		configDir, err := utils.GetConfigDir()
		if err != nil {
			log.Fatalln(err)
		}
		viper.AddConfigPath(configDir)

		viper.SetConfigName(DEFAULT_CONFIG_FILE_NAME)
		viper.SetConfigType("toml")
	}

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalln("Config file not found, please create one using 'config' command.")
	}
}
