package utils

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func LoadConfig() (*PMTConfig, error) {
	content, err := os.ReadFile("./pmt-config.json")
	if err != nil {
		return nil, err
	}

	var config *PMTConfig

	err = json.Unmarshal(content, &config)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return config, nil
}

func FileOrDirExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}

	return false, err
}

func GetPMTGlobalDir() (string, error) {
	fp, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	fp = filepath.Join(fp, ".pmt")

	isExist, err := FileOrDirExists(fp)
	if err != nil {
		return "", err
	}

	if !isExist {
		os.Mkdir(fp, os.ModePerm)
	}

	return fp, nil
}

func HandleCreateGoBackend(dir string) (string, error) {
	workCmd := fmt.Sprintf("work init %s", dir)
	args := strings.Split(workCmd, " ")

	cmd, err := exec.Command("go", args...).Output()
	if err != nil {
		log.Fatal(err.Error())
		return "", err
	}

	return string(cmd), nil
}
