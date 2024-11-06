package registry

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/NickBlakW/pmt/helpers/utils"
)

func HandleRemoveRegisteredProj(reg string) (string, error) {
	projects, err := GetRegisteredProjects()
	if err != nil {
		return "", err
	}

	if _, ok := projects[reg]; !ok {
		return "", fmt.Errorf("'%s' not found in registry", reg)
	}

	var toWrite []string
	for k, v := range projects {
		if reg == k {
			continue
		}
		var p = fmt.Sprintf("%s@%s", k, v)

		toWrite = append(toWrite, p)
	}

	regFile, err := utils.GetPMTGlobalDir()
	if err != nil {
		return "", err
	}
	regFile = filepath.Join(regFile, REGISTRY_FILE)

	file, err := os.Create(regFile)
	if err != nil {
		return "", err
	}

	defer file.Close()

	for _, s := range toWrite {
		file.WriteString(fmt.Sprintf("%s\n", s))
	}

	return fmt.Sprintf("Project '%s' removed from registry", reg), nil
}