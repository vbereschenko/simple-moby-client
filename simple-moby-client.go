package main

import (
	"context"
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/network"
	"github.com/docker/docker/client"
	"os"
)

func main() {
	c, err := client.NewEnvClient()
	if err != nil {
		fmt.Printf("failed initializing moby client: %e", err)
		return
	}
	defer c.Close()

	ctx := context.Background()
	// missing validation of args
	commandName := os.Args[1]

	switch commandName {
	case "list":
		containers, err := c.ContainerList(ctx, types.ContainerListOptions{})
		if err != nil {
			fmt.Printf("failed querying for container list: %e", err)
			return
		}
		for _, container := range containers {
			fmt.Println(container.ID, container.Status)
		}

	case "run":
		config := &container.Config{
			Cmd: os.Args[3:],
			Image: os.Args[2],
		}

		cont, err := c.ContainerCreate(ctx, config, &container.HostConfig{}, &network.NetworkingConfig{}, "")
		if err != nil {
			fmt.Printf("failed creating container: %e", err)
			return
		}

		if err = c.ContainerStart(ctx, cont.ID, types.ContainerStartOptions{}); err != nil {
			fmt.Printf("failed starting container: %e", err)
			return
		}

		fmt.Println(cont.ID)

	case "stop":
		if err = c.ContainerStop(ctx, os.Args[2], nil); err != nil {
			fmt.Printf("failed stopping container: %e", err)
			return
		}
		fmt.Println("stopped")
	}
}
