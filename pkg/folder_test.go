package pkg

import (
	"testing"
)

func TestCheckFolder(t *testing.T) {
	t.Run("Current working directory is a folder", func(t *testing.T) {
		findFolder, _ := checkIsFolder(".")

		got := findFolder
		want := true

		if got != want {
			t.Errorf("got %t, want %t", got, want)
		}
	})

	t.Run("File path does not exist", func(t *testing.T) {
		findFolder, _ := checkIsFolder("/some/path")

		got := findFolder
		want := false

		if got != want {
			t.Errorf("got %t, want %t", got, want)
		}
	})

	t.Run("Path is a file", func(t *testing.T) {
		findFolder, _ := checkIsFolder("./folder.go")

		got := findFolder
		want := false

		if got != want {
			t.Errorf("got %t, want %t", got, want)
		}
	})
}

func TestFindDominantLanguage(t *testing.T) {
	t.Run("Folder returns Go", func(t *testing.T) {
		dominantLanguage, _ := DirectoryWalkthrough(".")

		got := dominantLanguage
		want := "Go"

		if got != want {
			t.Errorf("got %s, want %s", got, want)
		}
	})
}
