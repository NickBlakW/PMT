package cfg_cmd

import (
	"fmt"

	"github.com/NickBlakW/pmt/helpers/config"
	"github.com/spf13/cobra"
)

var ModCfgAuthorOpt string
var ModCfgFrontendOpt string
var ModCfgBackendOpt string
var ModCfgScriptsOpt []string

var ConfigModifyCmd = &cobra.Command{
	Use: "modify",
	Short: "Modify the config file in this project",
	Aliases: []string{ "mod" },
	Run: runModifyConfigCmd,
}

func runModifyConfigCmd(cmd *cobra.Command, args []string) {
	out, err := config.ModifyConfig(ModCfgAuthorOpt, ModCfgScriptsOpt, ModCfgBackendOpt, ModCfgFrontendOpt)
	if err != nil {
		fmt.Println(out, err.Error())
		return
	}

	fmt.Println(out)
}

func ConfigModCfgCmd() {
	ConfigModifyCmd.Flags().StringArrayVarP(&ModCfgScriptsOpt, "script", "s", []string{}, "Adds script handlers")
	ConfigModifyCmd.Flags().StringVarP(&ModCfgAuthorOpt, "author", "a", "", "Adds author to project")
	ConfigModifyCmd.Flags().StringVarP(&ModCfgFrontendOpt, "frontend", "f", "", "Adds a defined frontend to project")
	ConfigModifyCmd.Flags().StringVarP(&ModCfgBackendOpt, "backend", "b", "", "Adds a defined backend to project")
}