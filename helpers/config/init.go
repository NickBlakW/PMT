package config

import (
	"maps"

	"github.com/NickBlakW/pmt/helpers/utils"
)

func InitPMT(opt []string, name string, author string, frontendDir string, backendDir string) (string, error) {
	var pmtConfig utils.PMTConfig

	pmtConfig.Name = name
	pmtConfig.Author = author
	pmtConfig.Frontend = frontendDir
	pmtConfig.Backend = backendDir

	var scripts = make(map[string]interface{})

	for _, lang := range opt {
		// Create a go.work file
		if lang == "go" && backendDir != "" {
			utils.HandleCreateGoBackend(backendDir)
		}

		switch (lang) {
		case "go", "maven":
			temp := CreateTemplate(backendDir, lang)
			maps.Copy(scripts, temp)
		case "npm", "yarn":
			temp := CreateTemplate(frontendDir, lang)
			maps.Copy(scripts, temp)
		}
	}

	pmtConfig.Scripts = scripts

	if err := WriteConfig(pmtConfig); err != nil {
		return "Unable to initialize PMT", err
	}

	return "PMT project initialized", nil
}
