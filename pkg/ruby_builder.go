package pkg

import (
	"log"
	"os"
)

type RubyBuilder struct {
}

func (g *RubyBuilder) buildImage(f *os.File) {
	buildRuby := `FROM ruby:latest
RUN mkdir /usr/src/app
ADD . /usr/src/app/
WORKDIR /usr/src/app/
CMD ["/usr/src/app/main.rb"]
`

	if _, err := f.Write([]byte(buildRuby)); err != nil {
		f.Close()
		log.Fatal(err)
	}

	if err := f.Close(); err != nil {
		log.Fatal(err)
	}
}
