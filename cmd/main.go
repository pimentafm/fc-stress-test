package main

import (
	"os"

	"github.com/pimentafm/fc-stress-test/pkg"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "stress-test",
	Short: "Stress testing utility",
	Long: `A tool designed to perform stress tests on web applications and APIs.
This utility allows you to simulate high levels of traffic to evaluate performance and reliability under load.
Key features include configurable request rates, concurrency levels, and detailed performance metrics.`,
	Run: func(cmd *cobra.Command, args []string) {
		url, _ := cmd.Flags().GetString("url")
		requests, _ := cmd.Flags().GetInt("requests")
		concurrency, _ := cmd.Flags().GetInt("concurrency")

		input := &pkg.Input{
			Url:      url,
			Requests: requests,
			Workers:  concurrency,
		}

		pkg.ExecuteTests(input)
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	rootCmd.Flags().StringP("url", "u", "", "URL")
	rootCmd.Flags().IntP("requests", "r", 100, "Number of requests")
	rootCmd.Flags().IntP("concurrency", "c", 5, "Number of concurrent requests")
	rootCmd.MarkFlagRequired("url")
	rootCmd.MarkFlagRequired("requests")
	rootCmd.MarkFlagRequired("concurrency")
}

func main() {
	Execute()
}
