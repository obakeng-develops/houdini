package pkg

import (
	"context"
	"fmt"
	"log/slog"
	"os"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"github.com/joho/godotenv"
)

type ImageBuild struct {
	tag      string
	language string
	user     string
	command  []string
}

func BuildImage(i *ImageBuild) {

}

func DockerExample() error {
	err := godotenv.Load()
	if err != nil {
		slog.Error("Error loading .env file", "err", err)
	}

	dockerSocketPath := os.Getenv("DOCKER_SOCKET_PATH")

	cli, err := client.NewClientWithOpts(client.WithHost(dockerSocketPath))
	if err != nil {
		panic(err)
	}

	images, err := cli.ImageList(context.TODO(), types.ImageListOptions{})
	if err != nil {
		panic(err)
	}

	for _, ctr := range images {
		fmt.Printf("%s %s\n", ctr.ID, ctr.Size)
	}

	return nil
}
