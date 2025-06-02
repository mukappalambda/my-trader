package cmd

import (
	"github.com/spf13/cobra"
)

var schemasCmd = &cobra.Command{
	Use:     "schemas",
	Aliases: []string{"schema", "sche", "sch"},
	Short:   "Communicate with the schema registry service",
	Long:    ``,
}

func init() {
	rootCmd.AddCommand(schemasCmd)
}
