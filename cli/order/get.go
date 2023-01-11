package order

import (
	"context"
	"fmt"

	"github.com/merext/binancectl/v2/api/order"
	"github.com/merext/binancectl/v2/utils"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var GetCmd = &cobra.Command{
	Use:   "get",
	Short: "Get order details",
	Run: func(cmd *cobra.Command, args []string) {
		symbol, _ := cmd.Flags().GetString("symbol")
		orderId, _ := cmd.Flags().GetInt64("orderId")
		order, err := order.Get(context.Background(), symbol, orderId)
		if err != nil {
			fmt.Print(err)
		}
		utils.OPrint(order)
	},
}

func init() {
	GetCmd.Flags().StringP("symbol", "s", "", "Symbol")
	GetCmd.MarkFlagRequired("symbol")

	GetCmd.Flags().Int64P("orderId", "i", 0, "Order ID")
	GetCmd.MarkFlagRequired("orderId")
}
