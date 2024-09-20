package helpers

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func AddDirOrFile(path string) (*os.File, error) {
	var _, err = os.Stat(path)

	if os.IsExist(err) {
		return nil, err
	}

	if err := os.MkdirAll(filepath.Dir(path), os.ModePerm); err != nil {
		return nil, err
	}
	return os.Create(path)
}

func HandleRemoveDirOrFile(path string) (string, error) {
	// Check if file is chosen to be deleted
	splitPath := strings.Split(path, "/")

	possibleFile := splitPath[len(splitPath)-1]
	hasFileExt := strings.Contains(possibleFile, ".")

	// handle if a file or file extension has been requested for deletion
	if hasFileExt {
		return handleFileDeletion(possibleFile, path, splitPath)
	}

	// if no file type has been requested delete the entire directory
	err := os.RemoveAll(path)
	if err != nil {
		return fmt.Sprintf("Unable to remove directory '%s'", path), err
	}

	return fmt.Sprintf("Deleted directory '%s' and everything in it", path), nil
}

func handleFileDeletion(filename string, path string, splitPath []string) (string, error) {
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
