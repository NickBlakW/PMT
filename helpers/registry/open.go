package registry

import (
	"os/exec"

	"github.com/NickBlakW/pmt/helpers/utils"
)

func OpenRegistry() error {
	path, err := utils.GetPMTGlobalDir()
	if err != nil {
		return nil
	}

	cmd := exec.Command("code", REGISTRY_FILE)	
	cmd.Dir = path

	err = cmd.Run()
	if err != nil {
		return err
	}

	return nil
}