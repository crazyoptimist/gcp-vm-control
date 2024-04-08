/*
Copyright © 2024 crazyoptimist <hey@crazyoptimist.net>
*/
package cmd

import (
	"bufio"
	"fmt"
	"gcp-vm-control/pkg/utils"
	"log"
	"os"
	"path/filepath"

	"github.com/pelletier/go-toml/v2"
	"github.com/spf13/cobra"
)

// configCmd represents the config command
var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Configure the application.",
	Long: `Provide necessary config values, then the application will
generate a configuration file.`,
	Run: func(cmd *cobra.Command, args []string) {
		config := readinConfig()

		configDir, err := utils.GetConfigDir()
		if err != nil {
			log.Fatalln(err)
		}

		err = generateConfigFile(*config, filepath.Join(configDir, DEFAULT_CONFIG_FILE_NAME)+".toml")
		if err != nil {
			log.Fatalln(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(configCmd)
}

type appConfig struct {
	ProjectID    string
	Zone         string
	InstanceName string
}

func readinConfig() *appConfig {
	stdinReader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter your Google Cloud ProjectID: ")
	ProjectID := getUserInput(stdinReader)
	fmt.Print("Enter your instance's Zone Name: ")
	Zone := getUserInput(stdinReader)
	fmt.Print("Enter your instance name: ")
	InstanceName := getUserInput(stdinReader)

	return &appConfig{
		ProjectID,
		Zone,
		InstanceName,
	}
}

func getUserInput(reader *bufio.Reader) string {
	val, err := reader.ReadString('\n')
	if err != nil {
		log.Fatalln("Malformed input. Please start over.")
	}
	// Remove newline character from the end of the input
	val = val[:len(val)-1]
	return val
}

func generateConfigFile(config appConfig, path string) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	err = toml.NewEncoder(file).Encode(config)
	if err != nil {
		return err
	}

	fmt.Printf("Config file created successfully: %s\n", path)
	return nil
}
