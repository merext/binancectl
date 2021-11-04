package order

import (
	"fmt"

	"github.com/merext/binancectl/v2/api/order"
	"github.com/merext/binancectl/v2/utils"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var SellCMD = &cobra.Command{
	Use:   "sell",
	Short: "Sell order",
	Run: func(cmd *cobra.Command, args []string) {
		symbol, _ := cmd.Flags().GetString("symbol")
		quantity, _ := cmd.Flags().GetString("quantity")
		price, _ := cmd.Flags().GetString("price")

		order, err := order.Sell(symbol, quantity, price)
		if err != nil {
			fmt.Print(err)
		}
		utils.OPrint(order)
	},
}

func init() {
	SellCMD.Flags().StringP("symbol", "s", "", "Symbol")
	SellCMD.MarkFlagRequired("symbol")

	SellCMD.Flags().StringP("quantity", "q", "", "Quantaty")
	SellCMD.MarkFlagRequired("quantity")

	SellCMD.Flags().StringP("price", "p", "", "Price")
	// SellCMD.MarkFlagRequired("price")
}
