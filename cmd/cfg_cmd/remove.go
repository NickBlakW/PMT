package cfg_cmd

import (
	"github.com/NickBlakW/pmt/helpers/config"
	"github.com/spf13/cobra"
)

var ConfigRemoveCmd = &cobra.Command{
	Use: "remove",
	Short: "Removes part of the projects config file",
	Aliases: []string{ "rem" },
	Run: runRemoveCmd,
}

func runRemoveCmd(cmd *cobra.Command, args []string) {
	config.RemoveConfigOption("")
}