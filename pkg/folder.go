package pkg

import (
	"os"
)

func checkIsFolder(path string) (bool, error) {
	folder, err := os.Stat(path)
	if err != nil {
		return false, err
	}

	if folder.IsDir() {
		return true, nil
	}

	return false, nil
}
