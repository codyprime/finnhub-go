package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	finnhub "github.com/Finnhub-Stock-API/finnhub-go/v2"
)

func main() {
	// Get API key from environment variable or use placeholder
	apiKey := os.Getenv("FINNHUB_API_KEY")
	if apiKey == "" {
		fmt.Println("Warning: FINNHUB_API_KEY environment variable not set. Using placeholder.")
		fmt.Println("Set FINNHUB_API_KEY environment variable to use real API calls.")
		apiKey = "<API_key>"
	}

	// Initialize configuration
	cfg := finnhub.NewConfiguration()
	cfg.AddDefaultHeader("X-Finnhub-Token", apiKey)
	finnhubClient := finnhub.NewAPIClient(cfg).DefaultApi

	ctx := context.Background()

	// Test 1: Quote
	// fmt.Println("\n=== Test 1: Quote ===")
	// quote, resp, err := finnhubClient.Quote(ctx).Symbol("AAPL").Execute()
	// if err != nil {
	// 	fmt.Printf("Error: %v\n", err)
	// 	if resp != nil {
	// 		fmt.Printf("Response status: %s\n", resp.Status)
	// 	}
	// } else {
	// 	// Use JSON marshaling to properly display pointer values
	// 	quoteJSON, _ := json.MarshalIndent(quote, "", "  ")
	// 	fmt.Printf("Quote:\n%s\n", string(quoteJSON))
	// }

	// // Test 2: Company Profile
	// fmt.Println("\n=== Test 2: Company Profile ===")
	// profile, resp, err := finnhubClient.CompanyProfile(ctx).Symbol("AAPL").Execute()
	// if err != nil {
	// 	fmt.Printf("Error: %v\n", err)
	// 	if resp != nil {
	// 		fmt.Printf("Response status: %s\n", resp.Status)
	// 	}
	// } else {
	// 	// Use JSON marshaling to properly display pointer values
	// 	profileJSON, _ := json.MarshalIndent(profile, "", "  ")
	// 	fmt.Printf("Company Profile:\n%s\n", string(profileJSON))
	// }

	// // Test 3: Company News
	// fmt.Println("\n=== Test 3: Company News ===")
	// news, resp, err := finnhubClient.CompanyNews(ctx).Symbol("AAPL").From("2024-01-01").To("2024-01-31").Execute()
	// if err != nil {
	// 	fmt.Printf("Error: %v\n", err)
	// 	if resp != nil {
	// 		fmt.Printf("Response status: %s\n", resp.Status)
	// 	}
	// } else {
	// 	fmt.Printf("Company News (showing first 3): %d articles found\n", len(news))
	// 	if len(news) > 0 {
	// 		for i := 0; i < len(news) && i < 3; i++ {
	// 			fmt.Printf("  - %s: %s\n", news[i].Headline, news[i].Url)
	// 		}
	// 	}
	// }

	// // Test 4: Stock Symbols
	// fmt.Println("\n=== Test 4: Stock Symbols ===")
	// stockSymbols, resp, err := finnhubClient.StockSymbols(ctx).Exchange("US").Execute()
	// if err != nil {
	// 	fmt.Printf("Error: %v\n", err)
	// 	if resp != nil {
	// 		fmt.Printf("Response status: %s\n", resp.Status)
	// 	}
	// } else {
	// 	fmt.Printf("Stock Symbols: %d symbols found\n", len(stockSymbols))
	// 	if len(stockSymbols) > 0 {
	// 		fmt.Printf("First 5 symbols:\n")
	// 		for i := 0; i < len(stockSymbols) && i < 5; i++ {
	// 			fmt.Printf("  - %s: %s\n", stockSymbols[i].Symbol, stockSymbols[i].Description)
	// 		}
	// 	}
	// }

	// // Test 5: Symbol Search
	// fmt.Println("\n=== Test 5: Symbol Search ===")
	// searchResult, resp, err := finnhubClient.SymbolSearch(ctx).Q("AAPL").Execute()
	// if err != nil {
	// 	fmt.Printf("Error: %v\n", err)
	// 	if resp != nil {
	// 		fmt.Printf("Response status: %s\n", resp.Status)
	// 	}
	// } else {
	// 	// Use JSON marshaling to properly display pointer values
	// 	searchJSON, _ := json.MarshalIndent(searchResult, "", "  ")
	// 	fmt.Printf("Search Results:\n%s\n", string(searchJSON))
	// }

	// // Test 6: Market News
	// fmt.Println("\n=== Test 6: Market News ===")
	// generalNews, resp, err := finnhubClient.MarketNews(ctx).Category("general").Execute()
	// if err != nil {
	// 	fmt.Printf("Error: %v\n", err)
	// 	if resp != nil {
	// 		fmt.Printf("Response status: %s\n", resp.Status)
	// 	}
	// } else {
	// 	fmt.Printf("Market News: %d articles found\n", len(generalNews))
	// 	if len(generalNews) > 0 {
	// 		fmt.Printf("First article: %s\n", generalNews[0].Headline)
	// 	}
	// }

	// Test 7: ETF Allocation by Symbol
	fmt.Println("\n=== Test 7: ETF Allocation by Symbol ===")
	allocation, resp, err := finnhubClient.EtfsAllocation(ctx).Symbol("SPY").Execute()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		if resp != nil {
			fmt.Printf("Response status: %s\n", resp.Status)
		}
	} else {
		allocationJSON, _ := json.MarshalIndent(allocation, "", "  ")
		fmt.Printf("ETF Allocation:\n%s\n", string(allocationJSON))
	}

	fmt.Println("\n=== Tests completed ===")
}
