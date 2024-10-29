package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/NickBlakW/pmt/helpers/utils"
	g "github.com/beetcb/ghdl"
	h "github.com/beetcb/ghdl/helper"
	"github.com/spf13/cobra"
)

var UpdateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update PMT to latest version",
	Run:   runUpdate,
}

func runUpdate(cmd *cobra.Command, args []string) {
	ghRelease := g.GHRelease{RepoPath: "NickBlakW/PMT"}
	ghReleaseDl, err := ghRelease.GetGHReleases(false)

	if err != nil {
		fmt.Println(fmt.Sprintf("get gh releases failed: %s", err), h.PrintModeErr)
		os.Exit(1)
	}
	ghReleaseDl.BinaryName = "pmt.exe"
	pmtPath, err := utils.GetPMTGlobalDir()
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println(fmt.Sprintf("start downloading %s", fmt.Sprint(filepath.Base(ghReleaseDl.Url), h.SprintOptions{PromptOff: true, PrintMode: h.PrintModeSuccess})), h.PrintModeInfo)
	if err := ghReleaseDl.DlTo(pmtPath); err != nil {
		h.Println(fmt.Sprintf("download failed: %s", err), h.PrintModeErr)
		os.Exit(1)
	}

	if err := ghReleaseDl.ExtractBinary(); err != nil {
		switch err {
		case g.ErrNeedInstall:
			h.Println(fmt.Sprintf("%s. You can install it with the appropriate commands", err), h.PrintModeInfo)
			os.Exit(0)
		case g.ErrNoBin:
			h.Println(fmt.Sprintf("%s. Try to specify binary name flag", err), h.PrintModeInfo)
			os.Exit(0)
		default:
			h.Println(fmt.Sprintf("extract failed: %s", err), h.PrintModeErr)
			os.Exit(1)
		}
	}
	h.Println(fmt.Sprintf("saved executable to %s", ghReleaseDl.BinaryName), h.PrintModeSuccess)
	if err := os.Chmod(ghReleaseDl.BinaryName, 0777); err != nil {
		h.Println(fmt.Sprintf("chmod failed: %s", err), h.PrintModeErr)
	}
}
