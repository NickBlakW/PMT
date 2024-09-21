package cmd

import (
	"fmt"

	"github.com/NickBlakW/pmt/helpers"
	"github.com/spf13/cobra"
)

var InitOption []string
var InitNameOpt string
var InitAuthorOpt string
var InitFrontendOpt string
var InitBackendOpt string

var InitCmd = &cobra.Command{
	Use: "initialize",
	Short: "Initialize a project with custom settings",
	Aliases: []string{"init"},
	// Args: cobra.ExactArgs(1),
	Run: runInit,
}

func runInit(cmd *cobra.Command, args []string) {
	var _, err = helpers.InitializePMT(
		InitOption,
		InitNameOpt,
		InitAuthorOpt,
		InitFrontendOpt,
		InitBackendOpt)
	if err != nil {
		fmt.Println("PMT project already initialized")
		return
	}

	fmt.Println("PMT project initialized")
}

func ConfigInitCmd() {
	InitCmd.Flags().StringArrayVarP(&InitOption, "with", "w", []string{}, "Initializes project with chosen option")
	InitCmd.Flags().StringVarP(&InitNameOpt, "name", "n", "<Project name here>", "Sets project name")
	InitCmd.Flags().StringVarP(&InitAuthorOpt, "author", "a", "<Your name here>", "Sets project author")
	InitCmd.Flags().StringVarP(&InitBackendOpt, "backend", "b", "", "Sets project backend dir")
	InitCmd.Flags().StringVarP(&InitFrontendOpt, "frontend", "f", "", "Sets project frontend dir")
}
