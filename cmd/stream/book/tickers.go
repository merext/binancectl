package book

import (
	"github.com/merext/binancectl/v2/api/stream"
	"github.com/spf13/cobra"
)

var TickersCmd = &cobra.Command{
	Use:   "tickers",
	Short: "Book tickers for all symbols",
	Run: func(cmd *cobra.Command, args []string) {
		stream.AllBookTickers()
	},
}

func init() {
}
