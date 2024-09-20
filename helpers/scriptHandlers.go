package helpers

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

func RunScript(command string) (string, error) {
	content, err := os.ReadFile("./PMTConfig.json")
	if err != nil {
		return "PMT not initialized", err
	}

	var payload map[string]interface{}
	err = json.Unmarshal(content, &payload)
	if err != nil {
		log.Fatal("Incorrectly formatted JSON file")
	}

	iScript := payload["scripts"].(map[string]interface{})[command]
	script := fmt.Sprintf("%s", iScript)

	// handle multiple commands
	multiCmd := strings.Split(script, "&&")
	if len(multiCmd) > 1 {
		output := ""

		for _, command := range multiCmd {
			command = strings.Trim(command, " ")
			out, err := handleScript(command)

			if err != nil {
				log.Fatal(err)
				return "", err
			}
			output += out
		}
		return output, nil
	}

	out, err := handleScript(script)
	if err != nil {
		log.Fatal(err)
		return "", err
	}

	return out, err
}

func handleScript(script string) (string, error) {
	scriptArgs := strings.Split(script, " ")
	cmdArgs := scriptArgs[1:]

	cmd, err := exec.Command(scriptArgs[0], cmdArgs...).Output()

	// err = cmd.Run()
	if err != nil {
		log.Fatalf("Script failed with error: %s\n", err.Error())
	}

	return string(cmd), nil
}