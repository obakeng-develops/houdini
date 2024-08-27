package pkg

import (
	"io/fs"
	"log/slog"
	"os"
	"path/filepath"

	"github.com/go-enry/go-enry/v2"
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

func DirectoryWalkthrough(path string) (string, error) {
	validFolder, err := checkIsFolder(path)
	if err != nil {
		return "", err
	}

	if validFolder {
		languageList := make(map[string]int)

		// Recursively walk through given directory path and determine language of each file
		err := filepath.Walk(path, func(filePath string, info fs.FileInfo, err error) error {
			if err != nil {
				return err
			}

			if info.Mode().IsRegular() {
				lang, _ := enry.GetLanguageByExtension(info.Name())

				if _, exists := languageList[lang]; exists {
					languageList[lang]++
				} else {
					languageList[lang] = 1
				}
			}

			return err
		})
		if err != nil {
			slog.Error("An error occurred while walking through the directory", "err", err)
		}
		return determineDominantLanguage(languageList), nil
	}

	return "", nil
}

func determineDominantLanguage(languageList map[string]int) string {
	maxKey := ""
	maxValue := 0

	// Find the key with the most number of files
	// [Go: 2, Ruby: 1, Txt: 1]
	// Should return Go
	for key, value := range languageList {
		if key != "" {
			if value > maxValue {
				maxKey = key
				maxValue = value
			}
		}
	}

	return maxKey
}
