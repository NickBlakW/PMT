package cmd

import (
	"github.com/NickBlakW/pmt/cmd/reg_cmd"
	"github.com/spf13/cobra"
)

var RegistryCmd = &cobra.Command{
	Use: "registry",
	Short: "Access the internal project registry in PMT",
	Aliases: []string{ "reg" },
	Args: cobra.ExactArgs(1),
}

func InitRegistryCmd() {
	reg_cmd.ConfigAddToRegCmd()
	RegistryCmd.AddCommand(reg_cmd.AddToRegCmd)

	reg_cmd.ConfigListRegCmd()
	RegistryCmd.AddCommand(reg_cmd.RegListCmd)

	RegistryCmd.AddCommand(reg_cmd.RegOpenCmd)
	
	RegistryCmd.AddCommand(reg_cmd.RemoveFromRegCmd)
}
