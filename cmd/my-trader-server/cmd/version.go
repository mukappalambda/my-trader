package cmd

import (
	"github.com/mukappalambda/my-trader/version"
	"github.com/spf13/cobra"
)

var name = `
                  __              __
  __ _  __ ______/ /________ ____/ /__ ___________ ___ _____  _____ ____
 /  ' \/ // /___/ __/ __/ _ ` + "`" + `/ _  / -_) __/___(_-</ -_) __/ |/ / -_) __/
/_/_/_/\_, /    \__/_/  \_,_/\_,_/\__/_/     /___/\__/_/  |___/\__/_/
      /___/

`

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "",
	RunE: func(cmd *cobra.Command, args []string) error {
		version.Version(name)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
