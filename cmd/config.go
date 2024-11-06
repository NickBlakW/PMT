package cmd

import (
	cfg "github.com/NickBlakW/pmt/cmd/cfg_cmd"
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
	// Run: runCfg,
}

func ConfigCfgCmd() {
	ConfigCmd.AddCommand(cfg.ConfigModifyCmd)
	cfg.ConfigModCfgCmd()

	ConfigCmd.AddCommand(cfg.DeleteConfigCmd)
	cfg.InitCfgDeleteCmd()

	ConfigCmd.AddCommand(cfg.ConfigRemoveCmd)
	// ConfigCmd.Flags().StringArrayVarP(&CfgScriptsOpt, "script", "s", []string{}, "Adds script handlers")
	// ConfigCmd.Flags().StringVarP(&CfgAuthorOpt, "author", "a", "", "Adds author to project")
	// ConfigCmd.Flags().StringVarP(&CfgFrontendOpt, "frontend", "f", "", "Adds a defined frontend to project")
	// ConfigCmd.Flags().StringVarP(&CfgBackendOpt, "backend", "b", "", "Adds a defined backend to project")
}