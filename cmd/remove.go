package cmd

import (
	"fmt"

	"github.com/NickBlakW/pmt/helpers"
	"github.com/spf13/cobra"
)

var RemoveCmd = &cobra.Command{
	Use: "remove",
	Short: "Remove or delete file(s) or directory",
	Aliases: []string{"rem", "delete", "del"},
	Args: cobra.ExactArgs(1),
	Run: runRemoveCmd,
}

func runRemoveCmd(cmd *cobra.Command, args []string) {
	res, err := helpers.HandleRemoveDirOrFile(args[0])
	if err != nil {
		fmt.Println(err)
		fmt.Println(res)
		return
	}

	fmt.Println(res)
}
