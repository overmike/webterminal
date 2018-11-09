package main

import (
	"context"
	"fmt"

	"github.com/docker/docker/client"
)

func main() {
	c, err := client.NewEnvClient()
	if err != nil {
		panic(c)
	}

	fmt.Printf("Client Version %v\n", c.ClientVersion())
	fmt.Printf("docker host %v\n", c.DaemonHost())

	v, err := c.ServerVersion(context.Background())
	if err != nil {
		panic(err)
	}
	fmt.Printf("Server Version %v\n", v)
}
