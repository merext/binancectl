package book

import (
	"github.com/spf13/cobra"
)

var BookCmd = &cobra.Command{
	Use:   "book",
	Short: "Book tickers",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
	BookCmd.AddCommand(TickerCmd)
	BookCmd.AddCommand(TickersCmd)
}
