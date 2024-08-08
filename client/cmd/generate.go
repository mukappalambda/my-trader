package cmd

import (
	"github.com/mukappalambda/my-trader/client/commands"
	"github.com/mukappalambda/my-trader/client/common"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var generateCmd = &cobra.Command{
	Use:     "generate",
	Aliases: []string{"gen"},
	Short:   "Generate a new schema",
	Long:    ``,
	Example: `
- my-trader-cli schemas genenrate
	`,
	Run: commands.RunGenerate,
}

func init() {
	generateCmd.Flags().StringP("output", "o", "", "output file name")
	err := viper.BindPFlag("output", generateCmd.Flags().Lookup("output"))
	common.PrintToStderrThenExit(err)
	schemasCmd.AddCommand(generateCmd)
}
