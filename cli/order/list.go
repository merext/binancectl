package order

import (
	"context"
	"fmt"

	"github.com/adshao/go-binance/v2"
	"github.com/merext/binancectl/v2/api/order"
	"github.com/merext/binancectl/v2/utils"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var ListCmd = &cobra.Command{
	Use:   "list",
	Short: "List orders",
	Run: func(cmd *cobra.Command, args []string) {
		symbol, _ := cmd.Flags().GetString("symbol")
		all, _ := cmd.Flags().GetBool("all")
		var orders []*binance.Order
		var err error
		if all == true {
			orders, err = order.List(context.Background(), symbol)
		} else {
			orders, err = order.ListOpen(context.Background(), symbol)
		}
		if err != nil {
			fmt.Print(err)
		}
		utils.OPrint(orders)
	},
}

func init() {
	ListCmd.Flags().StringP("symbol", "s", "", "Symbol")
	ListCmd.MarkFlagRequired("symbol")

	ListCmd.Flags().BoolP("all", "a", false, "All orders")
}
