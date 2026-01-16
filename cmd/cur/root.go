package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"currency-cli/internal/config"
	"currency-cli/internal/api"
	"github.com/atotto/clipboard"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "cur [amount] [from] [to]",
	Short: "A simple currency converter CLI",
	Long: `A simple currency converter CLI that allows you to convert between different currencies, list available currencies, and track currency pairs.

Examples:
  $ cur 100 USD EUR
  100 USD = 93.50 EUR

  $ cur list
  Available currencies: USD, EUR, GBP, JPY, ...

  $ cur track USD EUR
  Tracking USD->EUR: 1 USD = 0.935 EUR (last checked: 5 mins ago)
`,
	Args:  cobra.ExactArgs(3),
	Run: runConvertCommand,
}

func runConvertCommand(cmd *cobra.Command, args []string) {
	amount, err := strconv.ParseFloat(args[0], 64)
	if err != nil {
		log.Fatal("Invalid amount:", err)
	}

	fromCurrency := args[1]
	toCurrency := args[2]

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

	convertedAmount := amount * rate
	fmt.Printf("%.2f %s = %.2f %s\n", amount, strings.ToUpper(fromCurrency), convertedAmount, strings.ToUpper(toCurrency))

	if config.Cfg.CopyAnswer {
		convertedAmountStr := fmt.Sprintf("%.2f", convertedAmount)
		err := clipboard.WriteAll(convertedAmountStr)
		if err != nil {
			log.Fatal("Failed to copy to clipboard:", err)
		}
		fmt.Println("Converted amount copied to clipboard.")
	}
}
