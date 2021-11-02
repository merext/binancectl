package order

import (
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var OrderCmd = &cobra.Command{
	Use:   "order",
	Short: "Orders operations",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
	OrderCmd.AddCommand(GetCmd)
	OrderCmd.AddCommand(ListCmd)
	OrderCmd.AddCommand(CreateCmd)
}
