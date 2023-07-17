package container

import controller_v1 "auth-service/internal/controller/v1"

// ControllerProvider Inject Controller
func (c *Container) ControllerProvider() {
	// Controller
	if err := c.Container.Provide(controller_v1.NewController); err != nil {
		c.Error = err
	}

	if err := c.Container.Provide(controller_v1.NewPingController); err != nil {
		c.Error = err
	}
	if err := c.Container.Provide(controller_v1.NewAuthController); err != nil {
		c.Error = err
	}

}
