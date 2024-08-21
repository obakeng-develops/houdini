package pkg

import (
	"log"
	"os"
)

type ImageBuilder interface {
	buildImage(f *os.File)
}

type BuildOptions struct {
	Tag      string
	Language string
	Path     string
}

func (b *BuildOptions) GenerateDockerImage() error {
	f, err := os.OpenFile("Dockerfile", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}

	switch {
	case b.Language == "Go":
		b.buildImage(f)
	}

	return nil
}
