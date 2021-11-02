package order

import (
	"fmt"

	"github.com/merext/binancectl/v2/api/order"
	"github.com/merext/binancectl/v2/utils"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var CreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create order",
	Run: func(cmd *cobra.Command, args []string) {
		symbol, _ := cmd.Flags().GetString("symbol")
		quantity, _ := cmd.Flags().GetString("quantity")
		price, _ := cmd.Flags().GetString("price")

		order, err := order.Create(symbol, quantity, price)
		if err != nil {
			fmt.Print(err)
		}
		utils.OPrint(order)
	},
}

func init() {
	CreateCmd.Flags().StringP("symbol", "s", "", "Symbol")
	CreateCmd.MarkFlagRequired("symbol")

	CreateCmd.Flags().StringP("quantity", "q", "", "Quantaty")
	CreateCmd.MarkFlagRequired("quantity")

	CreateCmd.Flags().StringP("price", "p", "", "Price")
	CreateCmd.MarkFlagRequired("price")

}
