package server

import (
	controller_v1 "auth-service/internal/controller/v1"

	"github.com/gin-gonic/gin"
	"go.uber.org/dig"
)

// RouteRestfulService ...
type RouteRestfulService struct {
	dig.In

	GatewayV1 controller_v1.Gateway
}

func (c *Server) ConfigRouteRESTful(r *gin.Engine) {
	//g := r.Group("api/campaign-service")
	//v1 := g.Group("/v1")
	//
	//v1.GET("/ping", c.Gateway.GatewayV1.PingController.Ping)

	authGroup := r.Group("/auth")

	SetupAuthRouter(authGroup, c.Gateway.GatewayV1.AuthController)
}

func SetupAuthRouter(router *gin.RouterGroup, c *controller_v1.AuthController) {
	router.GET("/profile", c.GetProfile)
	router.PUT("/change-password", c.ChangePassword)
	router.PUT("/reset-password", c.ResetPassword)
}
