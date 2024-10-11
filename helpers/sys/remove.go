package sys

import (
	"fmt"
	"os"
	"strings"
)

func HandleRemoveDirOrFile(path string) (string, error) {
	// Check if file is chosen to be deleted
	splitPath := strings.Split(path, "/")

	possibleFile := splitPath[len(splitPath)-1]
	hasFileExt := strings.Contains(possibleFile, ".")

	// handle if a file or file extension has been requested for deletion
	if hasFileExt {
		return HandleFileDeletion(possibleFile, path, splitPath)
	}

	// if no file type has been requested delete the entire directory
	err := os.RemoveAll(path)
	if err != nil {
		return fmt.Sprintf("Unable to remove directory '%s'", path), err
	}

	return fmt.Sprintf("Deleted directory '%s' and everything in it", path), nil
}
