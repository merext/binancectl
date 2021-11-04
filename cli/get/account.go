package get

import (
	"context"
	"fmt"

	"github.com/merext/binancectl/v2/api"
	"github.com/merext/binancectl/v2/utils"
	"github.com/spf13/cobra"
)

// accountCmd represents the account command
var AccountCmd = &cobra.Command{
	Use:   "account",
	Short: "Get operations",
	Run: func(cmd *cobra.Command, args []string) {

		res, err := api.Client.NewGetAccountService().Do(context.Background())
		if err != nil {
			fmt.Println(err)
			return
		}
		utils.OPrint(res)
	},
}

func init() {

}
