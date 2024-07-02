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
