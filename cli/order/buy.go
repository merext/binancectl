package order

import (
	"context"
	"fmt"

	"github.com/merext/binancectl/v2/api/order"
	"github.com/merext/binancectl/v2/utils"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var BuyCMD = &cobra.Command{
	Use:   "buy",
	Short: "Buy order",
	Run: func(cmd *cobra.Command, args []string) {
		symbol, _ := cmd.Flags().GetString("symbol")
		quantity, _ := cmd.Flags().GetString("quantity")
		price, _ := cmd.Flags().GetString("price")

		order, err := order.Buy(context.Background(), symbol, quantity, price)
		if err != nil {
			fmt.Print(err)
		}
		utils.OPrint(order)
	},
}

func init() {
	BuyCMD.Flags().StringP("symbol", "s", "", "Symbol")
	BuyCMD.MarkFlagRequired("symbol")

	BuyCMD.Flags().StringP("quantity", "q", "", "Quantaty")
	BuyCMD.MarkFlagRequired("quantity")

	BuyCMD.Flags().StringP("price", "p", "", "Price")
	// BuyCMD.MarkFlagRequired("price")

}
