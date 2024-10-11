package registry

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/NickBlakW/pmt/helpers/utils"
)

func GetRegisteredProjects() (map[string]string, error) {
	dir, err := utils.GetPMTGlobalDir()
	if err != nil {
		return nil, err
	}

	registry := filepath.Join(dir, REGISTRY_FILE)

	content, err := os.ReadFile(registry)
	if err != nil {
		return nil, err
	}

	projLines := strings.Split(string(content), "\n")
	projects := make(map[string]string)

	for _, line := range projLines {
		if line == "" {
			continue
		}

		regPath := strings.Split(line, "@")
		projects[regPath[0]] = regPath[1]
	}

	return projects, nil
}