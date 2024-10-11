package reg_cmd

import (
	"fmt"

	"github.com/NickBlakW/pmt/helpers/registry"
	"github.com/spf13/cobra"
)

var RegOpenCmd = &cobra.Command{
	Use: "open",
	Short: "Opens the registry in VSCode",
	Aliases: []string{},
	Run: runOpenCmd,
}

func runOpenCmd(cmd *cobra.Command, args []string) {
	err := registry.OpenRegistry()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("Opening registry file")
}