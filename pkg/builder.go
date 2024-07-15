package pkg

type BuildOptions struct {
	Tag      string
	Language string
	Path     string
}

func (b *BuildOptions) GenerateImage() error {

	return nil
}
