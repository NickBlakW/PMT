package sys

import (
	"errors"
	"os"
	"path/filepath"

	"github.com/NickBlakW/pmt/helpers/utils"
)

func AddDirOrFile(path string) (*os.File, error) {
	// Check if file or directory exists
	exists, err := utils.FileOrDirExists(path)
	if err != nil {
		return nil, err
	}

	// abort if it already exists
	if exists {
		err = errors.New("file or directory already exists")
		return nil, err
	}
	
	// create dir or file and everything in it
	if pathErr := os.MkdirAll(filepath.Dir(path), os.ModePerm); pathErr != nil {
		return nil, pathErr
	}
	
	return os.Create(path)
}
