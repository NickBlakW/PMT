package helpers

import (
	"encoding/json"
	"errors"
	"log"
	"os"
	"os/exec"
	"strings"
)

func HandleOpenVSCode(path string) error {
	cmd := exec.Command("code", path)

	err := cmd.Run()
	if err != nil {
		return err
	}

	return nil
}

func GetProjectAuthors() (string, error) {
	content, err := os.ReadFile("PMTConfig.json")
	if err != nil {
		return "PMT not initialized", err
	}

	var payload map[string]interface{}
	err = json.Unmarshal(content, &payload)
	if err != nil {
		log.Fatal("Incorrectly formatted JSON file")
	}

	authorString := payload["author"]
	authors := strings.Split(authorString.(string), ",")

	for i, a := range authors {
		if strings.HasPrefix(a, "<") && strings.HasSuffix(a, ">") {
			err = errors.New("author not set for project")
			return "Please set author name or use 'PMT init' with the '-a' flag", err
		}

		authors[i] = strings.TrimSpace(authors[i])
	}

	return strings.Join(authors, "\n"), nil
}
