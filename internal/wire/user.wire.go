//go:build wireinject 
 // +build wireinject 
package wire

import (
	"github.com/google/wire"
	"github.com/kaito-coder/go-ecommerce-architecture/internal/controller"
	"github.com/kaito-coder/go-ecommerce-architecture/internal/repo"
	"github.com/kaito-coder/go-ecommerce-architecture/internal/service"
)

func InitUserRouterHandler() (*controller.UserController, error ) {
	wire.Build(
		repo.NewUserRepository,
		service.NewUserService,
		controller.NewUserController,
	)
	return new(controller.UserController), nil
}