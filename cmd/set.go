/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"gopkg.in/yaml.v2"
)

var PgConnection string
var CometBftEndpoint string
var ListenAddress string
var PollFrequency int
var MaxBlockPagination int
var MaxTxPagination int

// filesCmd represents the files command
var setsCmd = &cobra.Command{
	Use:   "sets",
	Short: "set necessary kwil indexer configs.",
	Long:  `This command will set the necessary configs to run the application.`,
	Run: func(cmd *cobra.Command, args []string) {
		writeConfig()
		displayConfigs()
	},
}

func init() {
	rootCmd.AddCommand(setsCmd)
	setsCmd.Flags().StringVarP(&PgConnection, "pg-connection", "c", "postgres://indexer:indexer123@localhost:5433/indexer?sslmode=disable", "Postgres connection string for the underlying database")
	viper.BindPFlag("pg-connection", setsCmd.PersistentFlags().Lookup("pg-connection"))

	setsCmd.Flags().StringVarP(&CometBftEndpoint, "cometbft-endpoint", "e", "https://localhost:26657", "CometBFT endpoint")
	viper.BindPFlag("cometbft-endpoint", setsCmd.PersistentFlags().Lookup("cometbft-endpoint"))

	setsCmd.Flags().StringVarP(&ListenAddress, "listen-address", "a", "8000", "Address on which the REST API will listen")
	viper.BindPFlag("listen-address", setsCmd.PersistentFlags().Lookup("listen-address"))

	setsCmd.Flags().IntVarP(&PollFrequency, "poll-frequency", "f", 5, "Frequency (in seconds) that the indexer should poll the node for changes")
	viper.BindPFlag("poll-frequency", setsCmd.PersistentFlags().Lookup("poll-frequency"))

	setsCmd.Flags().IntVarP(&MaxBlockPagination, "max-block-pagination", "b", 30, "Maximum number of records that can be included in a page from the /block endpoint")
	viper.BindPFlag("max-block-pagination", setsCmd.PersistentFlags().Lookup("max-block-pagination"))

	setsCmd.Flags().IntVarP(&MaxTxPagination, "max-tx-pagination", "x", 30, "Maximum number of records that can be included in a page from the /tx endpoint")
	viper.BindPFlag("max-tx-pagination", setsCmd.PersistentFlags().Lookup("max-tx-pagination"))
}

func writeConfig() {
	con := readFile()
	config := Config{
		PgConnection:       PgConnection,
		CometBftEndpoint:   CometBftEndpoint,
		ListenAddress:      ListenAddress,
		PollFrequency:      PollFrequency,
		MaxBlockPagination: MaxBlockPagination,
		MaxTxPagination:    MaxTxPagination,
	}
	if config.PgConnection == "" {
		config.PgConnection = con.PgConnection
	}

	if config.CometBftEndpoint == "" {
		config.CometBftEndpoint = con.CometBftEndpoint
	}

	if config.ListenAddress == "" {
		config.ListenAddress = con.ListenAddress
	}

	if config.PollFrequency == 0 {
		config.PollFrequency = con.PollFrequency
	}

	if config.MaxBlockPagination == 0 {
		config.MaxBlockPagination = con.MaxBlockPagination
	}

	if config.MaxTxPagination == 0 {
		config.MaxTxPagination = con.MaxTxPagination
	}
	data, err := yaml.Marshal(config)
	if err != nil {
		fmt.Printf("Error marshalling config to YAML: %v\n", err)
		return
	}

	// Check if the config directory exists, create it if not
	if _, err := os.Stat("config"); os.IsNotExist(err) {
		err := os.Mkdir("config", 0755)
		if err != nil {
			fmt.Printf("Error creating config directory: %v\n", err)
			return
		}
	}

	// Check if the config.yaml file exists, create it if not
	if _, err := os.Stat("config/config.yaml"); os.IsNotExist(err) {
		_, err := os.Create("config/config.yaml")
		if err != nil {
			fmt.Printf("Error creating config file: %v\n", err)
			return
		}
	}

	// Write configuration data to config/config.yaml
	err = os.WriteFile("config/config.yaml", data, 0644)
	if err != nil {
		fmt.Printf("Error writing config file: %v\n", err)
		return
	}

	fmt.Println("Config file written successfully")
}

type Config struct {
	PgConnection       string `yaml:"pg_connection"`
	CometBftEndpoint   string `yaml:"cometbft_endpoint"`
	ListenAddress      string `yaml:"listen_address"`
	PollFrequency      int    `yaml:"poll_frequency"`
	MaxBlockPagination int    `yaml:"max_block_pagination"`
	MaxTxPagination    int    `yaml:"max_tx_pagination"`
}
