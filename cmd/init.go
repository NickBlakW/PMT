package cmd

import (
	"fmt"

	"github.com/NickBlakW/pmt/helpers/config"
	"github.com/spf13/cobra"
)

var InitOption []string
var InitAuthorOpt string
var InitFrontendOpt string
var InitBackendOpt string

var InitCmd = &cobra.Command{
	Use: "initialize",
	Short: "Initialize a project with custom settings",
	Aliases: []string{"init"},
	Args: cobra.ExactArgs(1),
	Run: runInit,
}

func runInit(cmd *cobra.Command, args []string) {
	var _, err = config.InitPMT(
		InitOption,
		args[0],
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
	InitCmd.Flags().StringVarP(&InitAuthorOpt, "author", "a", "", "Sets project author")
	InitCmd.Flags().StringVarP(&InitBackendOpt, "backend", "b", "", "Sets project backend dir")
	InitCmd.Flags().StringVarP(&InitFrontendOpt, "frontend", "f", "", "Sets project frontend dir")

	InitCmd.MarkFlagRequired("name")
}
