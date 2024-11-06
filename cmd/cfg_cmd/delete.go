package cfg_cmd

import (
	"fmt"

	"github.com/NickBlakW/pmt/helpers/config"
	"github.com/spf13/cobra"
)

var RemoveFromRegOpt bool

var DeleteConfigCmd = &cobra.Command{
	Use: "delete",
	Short: "Deletes the config file for project",
	Aliases: []string{ "del" },
	Run: runDeleteCmd,
}

func runDeleteCmd(cmd *cobra.Command, args []string) {
	out, err := config.DeleteConfigFile(RemoveFromRegOpt)
	if err != nil {
		fmt.Println(out)
		return
	}

	fmt.Println(out)
}

func InitCfgDeleteCmd() {
	DeleteConfigCmd.Flags().BoolVarP(&RemoveFromRegOpt, "yes", "y", false, "Remove project from registry too")
}