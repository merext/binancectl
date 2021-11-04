package account

import (
	"github.com/spf13/cobra"
)

// GetCmd represents the get command
var AccountCmd = &cobra.Command{
	Use:   "account",
	Short: "Account data",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
	AccountCmd.AddCommand(InfoCmd)
}
