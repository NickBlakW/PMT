package helpers

import (
	"fmt"
	"os"
)

func InitializePMT(opt []string, name string, author string, frontendDir string, backendDir string) (*os.File, error) {
	file, err := os.Create("PMTConfig.json")

	if os.IsExist(err) {
		return nil, err
	}

	defer file.Close()

	var scripts string
	for i, rt := range opt {
		switch (rt) {
		case "go":
			scripts += withGo
		case "npm":
			scripts += withNPM
		case "yarn":
			scripts += withYarn
		case "maven":
			scripts += withMaven
		default:
			scripts += `"command": "<Your command here>"`
		}

		if len(opt) != i+1 {
			scripts += ","
		}
	}

	file.WriteString(formatFile(name, author, scripts, frontendDir, backendDir))

	return file, nil
}

func formatFile(name string, author string, scripts string, fDir string, bDir string) string {
	feDir := ""
	beDir := ""

	if fDir != "" {
		feDir = fmt.Sprintf(withFrontendDir, fDir)
	}
	if bDir != "" {
		beDir = fmt.Sprintf(withBackendDir, bDir)
	}

	return fmt.Sprintf(defaultFile, name, author, scripts, feDir, beDir)
}

var defaultFile = `{
	"name": "%s",
	"author": "%s",
	"scripts": {%s
	}%s%s
}`

var withGo = `
		"run": "go run .",
		"compile": "go build && echo 'Compiled successfully'",
		"clean": "go mod tidy && echo 'Cleaned project'"`

var withNPM = `
		"dev": "npm run dev",
		"build": "npm run build",
		"start": "npm run start"`

var withYarn = `
		"dev": "yarn dev",
		"build": "yarn build",
		"start": "yarn start"`

var withMaven = `
		"ci": "mvn clean install",
		"package": "mvn package",
		"run": "mvn clean compile exec:java"`

var withBackendDir = `,
	"backend": "%s"`

var withFrontendDir = `,
	"frontend": "%s"`
