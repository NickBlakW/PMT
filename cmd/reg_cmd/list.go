package reg_cmd

import (
	"fmt"

	"github.com/NickBlakW/pmt/helpers/registry"
	"github.com/spf13/cobra"
)

var RegVerboseFlag bool

var RegListCmd = &cobra.Command{
	Use: "list",
	Short: "Lists all registered projects on this computer",
	Aliases: []string{ "ls" },
	Run: runListCmd,
}

func runListCmd(cmd *cobra.Command, args []string) {
	out, err := registry.ListRegisteredProjects(RegVerboseFlag)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(out)
}

func ConfigListRegCmd() {
	RegListCmd.Flags().BoolVarP(&RegVerboseFlag, "verbose", "v", false, "Shows registered key and directory path of the registered projects")
}
