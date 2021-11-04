package account

import (
	"fmt"

	"github.com/merext/binancectl/v2/api/account"
	"github.com/merext/binancectl/v2/utils"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var InfoCmd = &cobra.Command{
	Use:   "info",
	Short: "Get account details",
	Run: func(cmd *cobra.Command, args []string) {
		res, err := account.Info()
		if err != nil {
			fmt.Print(err)
		}
		utils.OPrint(res)
	},
}

func init() {
}
