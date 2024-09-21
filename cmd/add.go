package cmd

import (
	"fmt"

	"github.com/NickBlakW/pmt/helpers"
	"github.com/spf13/cobra"
)

var AddCmd = &cobra.Command{
	Use: "add",
	Short: "Adds a file or directory to project",
	Args: cobra.ExactArgs(1),
	Run: runAdd,
}

func runAdd(cmd *cobra.Command, args []string) {
	var _, err = helpers.AddDirOrFile(args[0])

	if err != nil {
		if err.Error() == fmt.Sprintf("open %s: is a directory", args[0]) {
			fmt.Printf("Directory %s created!\n", args[0])
			return
		}

		fmt.Println(err.Error())
		return
	}

	fmt.Printf("%s created!\n", args[0])
}