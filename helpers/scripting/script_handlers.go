package scripting

import (
	"log"
	"os/exec"
	"strings"

	"github.com/NickBlakW/pmt/helpers/utils"
)

func RunScript(command string) (string, error) {
	var out string

	cfg, err := utils.LoadConfig()
	if err != nil {
		return "PMT not initialized", err
	}

	script := cfg.Scripts[command].(string)
	multiScript := strings.Split(script, "&&")
	if len(multiScript) > 1 {
		for _, cmd := range multiScript {
			cmd = strings.TrimSpace(cmd)
			res, err := handleScript(cmd)
			
			if err != nil {
				log.Fatal(err)
				return "", err
			}
			out += res
		}
		return out, nil
	}
	out, err = handleScript(script)
	if err != nil {
		log.Fatal(err)
		return "", err
	}

	return out, nil
}

func handleScript(script string) (string, error) {
	scriptArgs := strings.Split(script, " ")
	cmdArgs := scriptArgs[1:]

	cmd, err := exec.Command(scriptArgs[0], cmdArgs...).Output()

	if err != nil {
		log.Fatalf("Script failed with error: %s\n", err.Error())
	}

	return string(cmd), nil
}
