package helpers

import (
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func AddDirOrFile(path string) (*os.File, error) {
	var exists, err = fileOrDirExists(path)
	if err != nil {
		return nil, err
	}

	if exists {
		err = errors.New("file or directory already exists")
		return nil, err
	}
	
	if pathErr := os.MkdirAll(filepath.Dir(path), os.ModePerm); pathErr != nil {
		return nil, pathErr
	}
	
	return os.Create(path)
}

func fileOrDirExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}

	return false, err
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
