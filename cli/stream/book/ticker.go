package book

import (
	"github.com/merext/binancectl/v2/api/stream"
	"github.com/spf13/cobra"
)

var TickerCmd = &cobra.Command{
	Use:   "ticker",
	Short: "Book tickers for specific symbol",
	Run: func(cmd *cobra.Command, args []string) {
		symbol, _ := cmd.Flags().GetString("symbol")

		if symbol != "" {
			stream.BookTicker(symbol)
		}

		cmd.Help()
	},
}

func init() {
	TickerCmd.Flags().StringP("symbol", "s", "", "Symbol")
	TickerCmd.MarkFlagRequired("symbol")
}
