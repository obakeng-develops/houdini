package pkg

import (
	"log"
	"os"
)

type GoBuilder struct {
}

func (g *GoBuilder) buildImage(f *os.File) {
	buildGo := `# syntax=docker/dockerfile:1
	
FROM golang:1.19

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY * ./
`

	if _, err := f.Write([]byte(buildGo)); err != nil {
		f.Close()
		log.Fatal(err)
	}

	if err := f.Close(); err != nil {
		log.Fatal(err)
	}
}
