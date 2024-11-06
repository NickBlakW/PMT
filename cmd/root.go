/*
Copyright © 2024 NickBlakW
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var version = "0.1.3"

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:     "pmt",
	Version: version,
	Short:   "PMT is a project management tool",
	Long: `Inspired by the functionality of the package.json file from node and npm.
PMT can execute defined scripts for you.
It is also capable of both adding and removing both files and directories with simple commands.
`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.SetVersionTemplate(fmt.Sprintf("PMT version v%s\n", version))
	rootCmd.AddCommand(UpdateCmd)

	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	ConfigInitCmd()
	rootCmd.AddCommand(InitCmd)

	rootCmd.AddCommand(AddCmd)
	rootCmd.AddCommand(RemoveCmd)

	rootCmd.AddCommand(AuthorCmd)

	rootCmd.AddCommand(ExecCmd)
	rootCmd.AddCommand(CodeCommand)
	ConfigCodeCmd()

	rootCmd.AddCommand(RegistryCmd)
	InitRegistryCmd()

	rootCmd.AddCommand(ConfigCmd)
	ConfigCfgCmd()
}
