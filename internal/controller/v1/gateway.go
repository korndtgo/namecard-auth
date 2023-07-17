package controller_v1

import (
	"go.uber.org/dig"
)

// Controller ...
type Controller struct {
	Gateway Gateway
}

// Gateway ...
type Gateway struct {
	dig.In

	PingController *PingController
	AuthController *AuthController
}

// NewController ...
func NewController(g Gateway) *Controller {
	return &Controller{
		Gateway: g,
	}
}
