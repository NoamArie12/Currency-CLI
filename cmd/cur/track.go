package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"currency-cli/internal/api"
	"github.com/spf13/cobra"
)

var trackCmd = &cobra.Command{
	Use:   "track [from] [to]",
	Short: "Tracks a currency pair",
	Args:  cobra.ExactArgs(2),
	Run: runTrackCommand,
}

func runTrackCommand(cmd *cobra.Command, args []string) {
	fromCurrency := args[0]
	toCurrency := args[1]

	apiKey := os.Getenv("API_KEY")
	if apiKey == "" {
		log.Fatal("API_KEY not found in environment")
	}

	rates, err := api.GetExchangeRates(apiKey, fromCurrency)
	if err != nil {
		log.Fatal("Failed to get exchange rates:", err)
	}

	rate, ok := rates.ConversionRates[strings.ToUpper(toCurrency)]
	if !ok {
		log.Fatalf("Conversion rate for %s not found", toCurrency)
	}

	fmt.Printf("Tracking %s->%s: 1 %s = %.4f %s (last checked: %s)\n",
		strings.ToUpper(fromCurrency),
		strings.ToUpper(toCurrency),
		strings.ToUpper(fromCurrency),
		rate,
		strings.ToUpper(toCurrency),
		time.Now().Format("2006-01-02 15:04:05"))
}

func init() {
	rootCmd.AddCommand(trackCmd)
}

