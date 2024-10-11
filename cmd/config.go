package cmd

import (
	"fmt"

	"github.com/NickBlakW/pmt/helpers/config"
	"github.com/spf13/cobra"
)

var CfgAuthorOpt string
var CfgFrontendOpt string
var CfgBackendOpt string
var CfgScriptsOpt []string

var ConfigCmd = &cobra.Command{
	Use: "config",
	Short: "Handle the project configuration",
	Aliases: []string{ "cfg" },
	Args: cobra.ExactArgs(1),
	Run: runCfg,
}

func runCfg(cmd *cobra.Command, args []string) {
	inCmd := args[0]

	switch (inCmd) {
	case "modify", "mod":
		out, err := config.ModifyConfig(CfgAuthorOpt, CfgScriptsOpt, CfgBackendOpt, CfgFrontendOpt)
		if err != nil {
			fmt.Println(out, err.Error())
			return
		}

		fmt.Println(out)
	default:
		fmt.Println("No arguments given")
		return 
	}
}

func ConfigCfgCmd() {
	ConfigCmd.Flags().StringArrayVarP(&CfgScriptsOpt, "script", "s", []string{}, "Adds script handlers")
	ConfigCmd.Flags().StringVarP(&CfgAuthorOpt, "author", "a", "", "Adds author to project")
	ConfigCmd.Flags().StringVarP(&CfgFrontendOpt, "frontend", "f", "", "Adds a defined frontend to project")
	ConfigCmd.Flags().StringVarP(&CfgBackendOpt, "backend", "b", "", "Adds a defined backend to project")
}