package cmd

import (
	"github.com/mukappalambda/my-trader/client/commands"
	"github.com/mukappalambda/my-trader/client/common"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var getCmd = &cobra.Command{
	Use:     "get",
	Aliases: []string{"list"},
	Short:   "Retrieve the schemas",
	Long:    ``,
	Example: `
- my-trader-cli schemas get
- my-trader-cli schemas get --subject foo
- my-trader-cli schemas get --name bar
	`,
	Run: commands.RunGet,
}

func init() {
	getCmd.Flags().String("subject", "", "subject of schema")
	getCmd.Flags().StringP("name", "n", "", "name of schema")
	getCmd.Flags().String("schema-registry-url", "http://localhost:8081", "schema registry url")

	err := viper.BindPFlag("subject", getCmd.Flags().Lookup("subject"))
	common.PrintToStderrThenExit(err)
	err = viper.BindPFlag("name", getCmd.Flags().Lookup("name"))
	common.PrintToStderrThenExit(err)
	err = viper.BindPFlag("schema-registry-url", getCmd.Flags().Lookup("schema-registry-url"))
	common.PrintToStderrThenExit(err)
	schemasCmd.AddCommand(getCmd)
}
