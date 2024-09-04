package cmd

import (
	"github.com/mukappalambda/my-trader/client/commands"
	"github.com/mukappalambda/my-trader/client/common"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var sendCmd = &cobra.Command{
	Use:   "send",
	Short: "",
	Long:  ``,
	Run:   commands.RunSend,
}

func init() {
	sendCmd.Flags().String("server-url", "localhost:50051", "server url")
	sendCmd.Flags().StringP("topic", "t", "", "topic name")
	sendCmd.Flags().StringP("message", "m", "", "message")
	sendCmd.Flags().String("schema-registry-url", "http://localhost:8081", "schema registry url")
	sendCmd.Flags().String("schema", "", "schema name")
	err := viper.BindPFlag("server-url", sendCmd.Flags().Lookup("server-url"))
	common.PrintToStderrThenExit(err)
	err = viper.BindPFlag("topic", sendCmd.Flags().Lookup("topic"))
	common.PrintToStderrThenExit(err)
	err = viper.BindPFlag("message", sendCmd.Flags().Lookup("message"))
	common.PrintToStderrThenExit(err)
	err = viper.BindPFlag("schema-registry-url", sendCmd.Flags().Lookup("schema-registry-url"))
	common.PrintToStderrThenExit(err)
	err = viper.BindPFlag("schema", sendCmd.Flags().Lookup("schema"))
	common.PrintToStderrThenExit(err)
	rootCmd.AddCommand(sendCmd)
}
