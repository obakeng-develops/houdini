package pkg

import (
	"log"
	"os"
)

type GoBuilder struct {
}

func (g *GoBuilder) buildImage(f *os.File) {
	buildGo := `FROM golang:1.19

# Set destination for COPY
WORKDIR /app

# Download Go modules
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code. Note the slash at the end, as explained in
# https://docs.docker.com/reference/dockerfile/#copy
COPY *.go ./
`

	if _, err := f.Write([]byte(buildGo)); err != nil {
		f.Close()
		log.Fatal(err)
	}

	if err := f.Close(); err != nil {
		log.Fatal(err)
	}
}
