package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/kaito-coder/go-ecommerce-architecture/internal/service"
	"github.com/kaito-coder/go-ecommerce-architecture/pkg/response"
)

type UserController struct {
	userService *service.UserService
}

func NewUserController() *UserController {
	return &UserController{
		userService: service.NewUserService(),
	}
}

func (uc *UserController) GetUserById(c *gin.Context) {
	uid := c.Param("uid")
	data := uc.userService.GetUserById(uid)
	response.SuccessResponse(c, 20001, data)
}
