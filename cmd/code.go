package cmd

import (
	"fmt"

	"github.com/NickBlakW/pmt/helpers"
	"github.com/spf13/cobra"
)

var CodeCommand = &cobra.Command{
	Use: "code",
	Short: "Opens project (or directory) in VSCode",
	Aliases: []string{"open", "o"},
	Args: cobra.MinimumNArgs(0),
	Run: runCodeCmd,
}

func runCodeCmd(cmd *cobra.Command, args []string) {
	dir := args[0]

	err := helpers.HandleOpenVSCode(dir)
	if err != nil {
		fmt.Println("Failed to open directory")
		return
	}
	fmt.Println("Opening project...")
}
