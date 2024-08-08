package cmd

import (
	"github.com/spf13/cobra"
)

var schemasCmd = &cobra.Command{
	Use:     "schemas",
	Aliases: []string{"schema", "sche", "sch"},
	Short:   "A brief description of your command",
	Long:    ``,
}

func init() {
	rootCmd.AddCommand(schemasCmd)
}
