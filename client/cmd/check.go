package cmd

import (
	"github.com/mukappalambda/my-trader/client/commands"
	"github.com/mukappalambda/my-trader/client/common"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var checkCmd = &cobra.Command{
	Use:   "check",
	Short: "Validate a schema",
	Long:  ``,
	Run:   commands.RunCheck,
}

func init() {
	checkCmd.Flags().StringP("filename", "f", "", "filename (required)")
	_ = checkCmd.MarkFlagRequired("filename")
	err := viper.BindPFlag("filename", checkCmd.Flags().Lookup("filename"))
	common.PrintToStderrThenExit(err)
	schemasCmd.AddCommand(checkCmd)
}
