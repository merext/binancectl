package trade

import (
	"github.com/spf13/cobra"
)

var TradeCmd = &cobra.Command{
	Use:   "trade",
	Short: "Trades operations",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
	TradeCmd.AddCommand(GetCmd)
	TradeCmd.AddCommand(ListCmd)
}
