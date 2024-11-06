package config

import (
	"maps"

	"github.com/NickBlakW/pmt/helpers/utils"
)

func ModifyConfig(author string, scripts []string, backend string, frontend string) (string, error) {
	cfg, err := utils.LoadConfig()
	if err != nil {
		return "PMT not initialized", err
	}

	if author != "" { cfg.Author = author }
	if backend != "" { cfg.Backend = backend }
	if frontend != "" { cfg.Frontend = backend }

	if len(scripts) > 0 {
		scrs := make(map[string]interface{})

		if len(cfg.Scripts) != 0 { maps.Copy(scrs, cfg.Scripts) }

		for _, lang := range scripts {
			switch (lang) {
			case "go", "maven":
				temp := CreateTemplate("", lang)
				maps.Copy(scrs, temp)
			case "npm", "yarn":
				temp := CreateTemplate("", lang)
				maps.Copy(scrs, temp)
			}
		}

		cfg.Scripts = scrs
	}

	if err := WriteConfig(*cfg); err != nil {
		return "Error modifying configuration", err
	}

	return "Successfully updated config file", nil
}