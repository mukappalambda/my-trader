package cmd

import (
	"github.com/mukappalambda/my-trader/client/commands"
	"github.com/mukappalambda/my-trader/client/common"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var applyCmd = &cobra.Command{
	Use:     "apply",
	Aliases: []string{"add"},
	Short:   "Apply a configuration to a schema by file name or stdin",
	Long:    ``,
	Run:     commands.RunApply,
}

func init() {
	applyCmd.Flags().StringP("filename", "f", "", "filename (required)")
	_ = applyCmd.MarkFlagRequired("filename")
	applyCmd.Flags().String("schema-registry-url", "http://localhost:8081", "schema registry url")
	err := viper.BindPFlag("filename", applyCmd.Flags().Lookup("filename"))
	common.PrintToStderrThenExit(err)
	err = viper.BindPFlag("schema-registry-url", applyCmd.Flags().Lookup("schema-registry-url"))
	common.PrintToStderrThenExit(err)
	schemasCmd.AddCommand(applyCmd)
}
