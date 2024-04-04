/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"gopkg.in/yaml.v2"
)

var list bool

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "kwil-config",
	Short: "set necessary kwil indexer configs.",
	Long:  `This command will set the necessary configs to run the application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		if list {
			displayConfigs()
		}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
func init() {
	rootCmd.Flags().BoolVarP(&list, "list", "l", false, "Display more verbose output in console output. (default: false)")
	viper.BindPFlag("list", rootCmd.PersistentFlags().Lookup("list"))
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
func readFile() Config {
	data, err := os.ReadFile("config/config.yaml")
	if err != nil {
		log.Errorf("Error reading config file: %v", err)
		return Config{}
	}

	// Unmarshal the config data
	var config Config
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		log.Errorf("Error unmarshalling config data: %v", err)
		return Config{}
	}
	return config
}
func displayConfigs() {
	// Read the config file
	config := readFile()
	// Print the config to console
	fields := log.Fields{
		"PgConnection":       config.PgConnection,
		"CometBftEndpoint":   config.CometBftEndpoint,
		"ListenAddress":      config.ListenAddress,
		"PollFrequency":      config.PollFrequency,
		"MaxBlockPagination": config.MaxBlockPagination,
		"MaxTxPagination":    config.MaxTxPagination,
	}

	// Iterate over each field and log it on a separate line
	log.Info("Current Config")
	for key, value := range fields {
		log.WithField(key, value).Info("")
	}
}
