package sys

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func HandleFileDeletion(filename string, path string, splitPath []string) (string, error) {
	fileInfo := strings.Split(filename, ".")
	requestedExt := fileInfo[len(fileInfo)-1]

	dirPath := splitPath[:len(splitPath)-1]
	joinedDirPath := strings.Join(dirPath[:], "/")
	dir, err := os.ReadDir(joinedDirPath)

	if err != nil {
		log.Fatal("Path not valid")
	}

	// Handle wildcard with extension
	if fileInfo[0] == "*" {
		for _, f := range dir {
			if strings.Contains(f.Name(), requestedExt) {
				err = os.RemoveAll(fmt.Sprintf("%s/%s", joinedDirPath, f.Name()))
			}
		}
		if err != nil {
			return "Unable to remove file(s)", err
		}

		return fmt.Sprintf("Deleted '.%s' file(s) in directory '%s'", fileInfo[len(fileInfo)-1], joinedDirPath), nil
	}

	err = os.RemoveAll(path)
	if err != nil {
		return "Unable to remove file(s)", err
	}

	return fmt.Sprintf("Deleted '.%s' file(s) in directory '%s'", fileInfo[len(fileInfo)-1], joinedDirPath), nil
}