package helpers

import (
	"os/exec"

	"github.com/NickBlakW/pmt/helpers/registry"
)

func HandleOpenVSCode(reg string, pathFlag string) error {
	projects, err := registry.GetRegisteredProjects()
	if err != nil {
		return err
	}

	path := projects[reg]

	if path == "" {
		path = "."
	}
	if pathFlag != "" {
		path = pathFlag
	}

	cmd := exec.Command("code", ".")
	cmd.Dir = path

	err = cmd.Run()
	if err != nil {
		return err
	}

	return nil
}
