package main

import (
	"fmt"
	"log"
	"os"
	"sort"

	"currency-cli/internal/api"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists available currencies",
	Run: runListCommand,
}

func runListCommand(cmd *cobra.Command, args []string) {
	apiKey := os.Getenv("API_KEY")
	if apiKey == "" {
		log.Fatal("API_KEY not found in environment")
	}

	rates, err := api.GetExchangeRates(apiKey, "USD")
	if err != nil {
		log.Fatal("Failed to get exchange rates:", err)
	}

	var currencies []string
	for currency := range rates.ConversionRates {
		currencies = append(currencies, currency)
	}
	sort.Strings(currencies)

	fmt.Println("Available currencies:")
	for _, currency := range currencies {
		fmt.Println(currency)
	}
}

func init() {
	rootCmd.AddCommand(listCmd)
}
