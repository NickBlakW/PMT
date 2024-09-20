package cmd

import (
	"fmt"

	"github.com/NickBlakW/pmt/helpers"
	"github.com/spf13/cobra"
)

var ExecCmd = &cobra.Command{
	Use: "exec",
	Short: "Execute a cli script from the project",
	Aliases: []string{"run", "do"},
	Args: cobra.ExactArgs(1),
	Run: runExec,
}

func runExec(cmd *cobra.Command, args []string) {
	output, err := helpers.RunScript(args[0])
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(output)
}

func ConfigExecCmd() {
	
}
