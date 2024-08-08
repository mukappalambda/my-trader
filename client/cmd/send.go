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
	err := viper.BindPFlag("server-url", sendCmd.Flags().Lookup("server-url"))
	common.PrintToStderrThenExit(err)
	err = viper.BindPFlag("topic", sendCmd.Flags().Lookup("topic"))
	common.PrintToStderrThenExit(err)
	err = viper.BindPFlag("message", sendCmd.Flags().Lookup("message"))
	common.PrintToStderrThenExit(err)
	rootCmd.AddCommand(sendCmd)
}
