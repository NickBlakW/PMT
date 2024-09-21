/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "pmt",
	Short: "PMT is a project management tool",
	Long: `Inspired by the functionality of the package.json file from node and npm.
PMT can execute defined scripts for you.
It is also capable of both adding and removing both files and directories with simple commands.
`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.pmt.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	ConfigInitCmd()
	rootCmd.AddCommand(InitCmd)
	
	rootCmd.AddCommand(AddCmd)
	rootCmd.AddCommand(RemoveCmd)

	rootCmd.AddCommand(AuthorCmd)

	rootCmd.AddCommand(ExecCmd)
	rootCmd.AddCommand(CodeCommand)
}


