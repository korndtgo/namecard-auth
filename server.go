package main

import "auth-service/internal/container"

func main() {
	c := container.NewContainer()

	if err := c.Run().Error; err != nil {
		panic(err)
	}
}
