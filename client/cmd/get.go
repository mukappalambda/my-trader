package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var getCmd = &cobra.Command{
	Use:     "get",
	Aliases: []string{"list"},
	Short:   "Retrieve the schemas",
	Long:    ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("schema get called")
	},
}

func init() {
	schemasCmd.AddCommand(getCmd)
}
