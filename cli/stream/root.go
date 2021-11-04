package stream

import (
	"github.com/merext/binancectl/v2/cli/stream/book"
	"github.com/spf13/cobra"
)

// StreamCmd represents the get command
var StreamCmd = &cobra.Command{
	Use:   "stream",
	Short: "Stream data",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
	StreamCmd.AddCommand(book.BookCmd)

}
