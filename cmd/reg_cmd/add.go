package reg_cmd

import (
	"fmt"

	"github.com/NickBlakW/pmt/helpers/registry"
	"github.com/spf13/cobra"
)

var RegisterAltFlag string

var AddToRegCmd = &cobra.Command{
	Use: "add",
	Short: "Adds the current project to the registry",
	Run: runRegistryAdd,
}

func runRegistryAdd(cmd *cobra.Command, args[] string) {
	out, err := registry.RegisterProject(RegisterAltFlag)
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println(out)
}

func ConfigAddToRegCmd() {
	AddToRegCmd.Flags().StringVarP(&RegisterAltFlag, "alt", "a", "", "Registers project under alternate name")
}