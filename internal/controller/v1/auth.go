package controller_v1

import (
	"auth-service/internal/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

type AuthController struct {
}

func (c *AuthController) GetProfile(ctx *gin.Context) {
	//var req model.ListIdpRequestQuery
	//
	//if err := ctx.BindQuery(&req); err != nil {
	//	_ = ctx.Error(err)
	//
	//	ctx.JSON(http.StatusBadRequest, model.ErrorResponse{
	//		Code:    "400",
	//		Message: err.Error(),
	//	})
	//
	//	return
	//}
	//
	//resp, err := c.ndidService.ListIdPs(nil, req)
	//if err != nil {
	//	_ = ctx.Error(err)
	//
	//	ctx.JSON(http.StatusInternalServerError, model.ErrorResponse{
	//		Code:    "500",
	//		Message: err.Error(),
	//	})
	//
	//	return
	//}
	//
	ctx.JSON(http.StatusOK, model.GetProfileDto{})
}
func (c *AuthController) ChangePassword(ctx *gin.Context) {

	ctx.JSON(http.StatusOK, "")
}
func (c *AuthController) ResetPassword(ctx *gin.Context) {

	ctx.JSON(http.StatusOK, "")
}

// NewAuthController ...
func NewAuthController() *AuthController {
	return &AuthController{}
}
