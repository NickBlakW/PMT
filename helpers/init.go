package helpers

import (
	"fmt"
	"log"
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
			if backendDir != "" {
				withGo = fmt.Sprintf(withGo, backendDir)
				_, err := HandleCreateGoBackend(backendDir)
				if err != nil {
					log.Fatal(err.Error())
				}
			} else {
				withGo = fmt.Sprintf(withGo, ".", ".")
			}

			scripts += withGo
		case "npm":
			if frontendDir != "" {
				prefix := fmt.Sprintf(" --prefix %s", frontendDir)
				withNPM = fmt.Sprintf(withNPM, prefix)
			} else {
				withNPM = fmt.Sprintf(withNPM, "")
			}

			scripts += withNPM
		case "yarn":
			if frontendDir != "" {
				prefix := fmt.Sprintf(" --cwd %s", frontendDir)
				withYarn = fmt.Sprintf(withYarn, prefix)
			} else {
				withYarn = fmt.Sprintf(withYarn, "")
			}

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
		"run": "go run %[1]s",
		"compile": "go build %[1]s && echo 'Compiled successfully'"`

var withNPM = `
		"dev": "npm run dev%[1]s",
		"build": "npm run build%[1]s",
		"start": "npm run start%[1]s"`

var withYarn = `
		"dev": "yarn%[1]s dev",
		"build": "yarn%[1]s build",
		"start": "yarn%[1]s start"`

var withMaven = `
		"ci": "mvn clean install",
		"package": "mvn package",
		"run": "mvn clean compile exec:java"`

var withBackendDir = `,
	"backend": "%s"`

var withFrontendDir = `,
	"frontend": "%s"`
