package registry

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/NickBlakW/pmt/helpers/utils"
)

func RegisterProject(alt string) (string, error) {
	registry, err := utils.GetPMTGlobalDir()
	if err != nil {
		return "", err
	}

	registry = filepath.Join(registry, REGISTRY_FILE)

	file, err := os.OpenFile(registry, os.O_APPEND|os.O_WRONLY|os.O_CREATE, os.ModePerm)
	if err != nil {
		return "", err
	}
	
	defer file.Close()

	reg, err := getProjectName()
	if err != nil {
		log.Fatal(err)
		return "", err
	}

	if alt != "" { reg = alt }

	if isRegistered(reg) {
		return fmt.Sprintf("%s already registered", reg), nil
	}
	
	path, err := getCurrentDirPath()
	if err != nil {
		return "", err
	}

	if _, err = file.WriteString(fmt.Sprintf("%s@%s\n", reg, path)); err != nil {
		panic(err)
	}

	return fmt.Sprintf("%s registered", reg), nil
}

func isRegistered(reg string) bool {
	projects, err := GetRegisteredProjects()
	if err != nil {
		log.Fatal(err)

		return false
	}

	var _, ok = projects[reg]

	return ok
}

func getCurrentDirPath() (string, error) {
	dir, err := os.Getwd()
	if err != nil {
		return "", err
	}

	return dir, nil
}

func getProjectName() (string, error) {
	cfg, err := utils.LoadConfig()
	if err != nil {
		return "", err
	}

	return cfg.Name, nil
}
