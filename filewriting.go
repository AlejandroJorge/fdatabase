package fdatabase

import (
	"os"
)

func folderExists(path string) bool {
	fileInfo, err := os.Stat(path)
	if err != nil {
		return false
	}

	return fileInfo.IsDir()
}

func handleMissingDir(path string) error {
	if folderExists(path) {
		return nil
	}

	if !folderExists(path) {
		err := os.Mkdir(path, os.ModePerm)
		if err != nil {
			return err
		}
	}
	return nil
}
