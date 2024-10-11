package reg_cmd

import (
	"fmt"

	"github.com/NickBlakW/pmt/helpers/registry"
	"github.com/spf13/cobra"
)

var RemoveFromRegCmd = &cobra.Command{
	Use: "remove",
	Short: "Removes the current project from the registry",
	Aliases: []string{ "rem", "rm" },
	Args: cobra.ExactArgs(1),
	Run: runRegistryRemove,
}

func runRegistryRemove(cmd *cobra.Command, args []string) {
	out, err := registry.HandleRemoveRegisteredProj(args[0])
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(out)
}