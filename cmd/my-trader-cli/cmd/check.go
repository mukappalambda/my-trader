package cmd

import (
	"github.com/mukappalambda/my-trader/cmd/my-trader-cli/commands"
	"github.com/mukappalambda/my-trader/cmd/my-trader-cli/common"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var checkCmd = &cobra.Command{
	Use:   "check",
	Short: "Validate a schema",
	Long:  ``,
	Example: `
- my-trader-cli schemas check -f example-schema.json
	`,
	RunE: commands.RunCheck,
}

func init() {
	checkCmd.Flags().StringP("filename", "f", "", "filename (required)")
	_ = checkCmd.MarkFlagRequired("filename")
	err := viper.BindPFlag("filename", checkCmd.Flags().Lookup("filename"))
	common.PrintToStderrThenExit(err)
	schemasCmd.AddCommand(checkCmd)
}
