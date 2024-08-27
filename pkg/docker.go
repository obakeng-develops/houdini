package pkg

import (
	"os"

	"github.com/docker/docker/client"
)

func ListContainers() {
	dockerSocketPath := os.Getenv("DOCKER_SOCKET_PATH")

	apiClient, err := client.NewClientWithOpts(client.WithHost(dockerSocketPath))
	if err != nil {
		panic(err)
	}
	defer apiClient.Close()
}
