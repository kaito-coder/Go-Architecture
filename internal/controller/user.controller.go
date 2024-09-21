package controller

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/kaito-coder/go-ecommerce-architecture/internal/service"
	"github.com/kaito-coder/go-ecommerce-architecture/internal/vo"
	"github.com/kaito-coder/go-ecommerce-architecture/pkg/response"
)

type UserController struct {
	userService service.IUserService
}

func NewUserController(
	userService service.IUserService,
) *UserController {
	return &UserController{
		userService: userService,
	}
}

func (uc *UserController) Register(c *gin.Context) {
	var params vo.UserRegisterRequest
	if err := c.ShouldBindJSON(&params); err != nil {
		response.FailResponse(c, response.ErrCodeParamError, err.Error())
		return
	}
	fmt.Printf("params email: %v", params.Email)
	result := uc.userService.Register(params.Email, params.Purpose)
	response.SuccessResponse(c, result, nil)
}

// func (uc *UserController) GetUserById(c *gin.Context) {
// 	uid := c.Param("uid")
// 	data := uc.userService.GetUserById(uid)
// 	response.SuccessResponse(c, 20001, data)
// }
