package pkg

import (
	"log"
	"os"
)

type RubyBuilder struct {
}

func (g *RubyBuilder) buildImage(f *os.File) {
	buildRuby := `# syntax=docker/dockerfile:1
	
FROM ruby:latest
WORKDIR /usr/src/app/
COPY . ./usr/src/app
`

	if _, err := f.Write([]byte(buildRuby)); err != nil {
		f.Close()
		log.Fatal(err)
	}

	if err := f.Close(); err != nil {
		log.Fatal(err)
	}
}
