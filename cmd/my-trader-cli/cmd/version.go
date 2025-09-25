package cmd

import (
	"github.com/mukappalambda/my-trader/cmd/my-trader-cli/commands"
	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:     "version",
	Aliases: []string{"v"},
	Short:   "version",
	Long:    ``,
	Example: `
- my-trader-cli version
	`,
	RunE: commands.RunVersion,
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
