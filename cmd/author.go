package cmd

import (
	"fmt"

	"github.com/NickBlakW/pmt/helpers"
	"github.com/spf13/cobra"
)

var AuthorCmd = &cobra.Command{
	Use: "author",
	Short: "Prints the author name(s)",
	Aliases: []string{"a"},
	Run: runAuthorCmd,
}

func runAuthorCmd(cmd *cobra.Command, args []string) {
	authors, err := helpers.GetProjectAuthors()
	if err != nil {
		fmt.Println(err)
		fmt.Println(authors)
		return
	}

	fmt.Println("Author(s) of this project:\n" + authors)
}
