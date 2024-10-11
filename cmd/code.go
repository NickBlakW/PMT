package cmd

import (
	"fmt"

	"github.com/NickBlakW/pmt/helpers"
	"github.com/spf13/cobra"
)

var CodeCmdPath string

var CodeCommand = &cobra.Command{
	Use: "code",
	Short: "Opens project (or directory) in VSCode",
	Aliases: []string{"open", "o"},
	Run: runCodeCmd,
}

func runCodeCmd(cmd *cobra.Command, args []string) {
	var reg = ""

	if len(args) >= 1 {
		reg = args[0]
	}

	err := helpers.HandleOpenVSCode(reg, CodeCmdPath)
	if err != nil {
		fmt.Println("Failed to open directory")
		return
	}
	fmt.Println("Opening project...")
}

func ConfigCodeCmd() {
	CodeCommand.Flags().StringVarP(&CodeCmdPath, "path", "p", "", "Open specific path from this directory")
}
