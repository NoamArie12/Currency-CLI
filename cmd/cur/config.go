package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"log"
	"os"
	"os/user"
	"path/filepath"
)

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Manage application configuration",
}

var setCmd = &cobra.Command{
	Use:   "set [key] [value]",
	Short: "Set a configuration value",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		key := args[0]
		value := args[1]

		viper.Set(key, value)

		usr, err := user.Current()
		if err != nil {
			log.Fatal(err)
		}
		configPath := filepath.Join(usr.HomeDir, ".config", "currency-cli")
		configFilePath := filepath.Join(configPath, "config.yaml")

		if err := os.MkdirAll(configPath, os.ModePerm); err != nil {
			log.Fatalf("Error creating config directory: %s", err)
		}

		if err := viper.WriteConfigAs(configFilePath); err != nil {
			log.Fatalf("Error writing config file: %s", err)
		}
		fmt.Printf("Set %s to %s\n", key, value)
	},
}

func init() {
	configCmd.AddCommand(setCmd)
	rootCmd.AddCommand(configCmd)
}
 