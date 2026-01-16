package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

// ExchangeRateResponse defines the structure for the API response.
type ExchangeRateResponse struct {
	Result          string             `json:"result"`
	BaseCode        string             `json:"base_code"`
	ConversionRates map[string]float64 `json:"conversion_rates"`
}

// GetExchangeRates retrieves exchange rates from the API.
func GetExchangeRates(apiKey, base string) (ExchangeRateResponse, error) {
	url := fmt.Sprintf("https://v6.exchangerate-api.com/v6/%s/latest/%s", apiKey, base)
	
	// Create a new HTTP client with a timeout
	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Get(url)
	if err != nil {
		return ExchangeRateResponse{}, fmt.Errorf("failed to make request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		body, _ := io.ReadAll(resp.Body)
		return ExchangeRateResponse{}, fmt.Errorf("request failed with status %d: %s", resp.StatusCode, string(body))
	}

	var data ExchangeRateResponse
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return ExchangeRateResponse{}, fmt.Errorf("failed to decode API response: %w", err)
	}

	return data, nil
}
