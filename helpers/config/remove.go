package config

import (
	"fmt"
	"os"

	"github.com/NickBlakW/pmt/helpers/registry"
	"github.com/NickBlakW/pmt/helpers/utils"
)

func RemoveConfigOption(cfgPart string) (string, error) {
	cfg, err := utils.LoadConfig()
	if err != nil {
		return "", err
	}

	switch (cfgPart) {
	case "author":
		var tmp string
		cfg.Author = tmp
	}
	WriteConfig(*cfg)

	return "", nil
}

func DeleteConfigFile(shouldRemove bool) (string, error) {
	if shouldRemove {
		cfg, err := utils.LoadConfig()
		if err != nil {
			return "Unable to load config file", err
		}

		out, err := registry.HandleRemoveRegisteredProj(cfg.Name)
		if err != nil {
			return "Error removing project from registry", err
		}

		fmt.Println(out)
	}

	err := os.Remove("./pmt-config.json")
	if err != nil {
		return "Config file not present in current directory", err
	}

	return "Removed config file from project", nil
}