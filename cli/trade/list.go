package trade

import (
	"fmt"

	"github.com/merext/binancectl/v2/api/trade"
	"github.com/merext/binancectl/v2/utils"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var ListCmd = &cobra.Command{
	Use:   "list",
	Short: "List trades",
	Run: func(cmd *cobra.Command, args []string) {
		symbol, _ := cmd.Flags().GetString("symbol")
		prices, err := trade.List(symbol)
		if err != nil {
			fmt.Print(err)
		}
		utils.OPrint(prices)
	},
}

func init() {
	ListCmd.Flags().StringP("symbol", "s", "", "Symbol")
	ListCmd.MarkFlagRequired("symbol")
}
