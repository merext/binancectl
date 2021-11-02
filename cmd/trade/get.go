package trade

import (
	"fmt"

	"github.com/merext/binancectl/v2/api/trade"
	"github.com/merext/binancectl/v2/utils"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var GetCmd = &cobra.Command{
	Use:   "get",
	Short: "Get trade details",
	Run: func(cmd *cobra.Command, args []string) {
		symbol, _ := cmd.Flags().GetString("symbol")
		order, err := trade.Get(symbol)
		if err != nil {
			fmt.Print(err)
		}
		utils.OPrint(order)
	},
}

func init() {
	GetCmd.Flags().StringP("symbol", "s", "", "Symbol")
	GetCmd.MarkFlagRequired("symbol")
}
