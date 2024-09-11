package user

import (
	"github.com/gin-gonic/gin"
	"github.com/kaito-coder/go-ecommerce-architecture/internal/controller"
	"github.com/kaito-coder/go-ecommerce-architecture/internal/repo"
	"github.com/kaito-coder/go-ecommerce-architecture/internal/service"
)

	type UserRouter struct {}

	func (ur *UserRouter) InitUserRouter(Router *gin.RouterGroup) {
		// public router 
		// non dependency injection
		uR := repo.NewUserRepository()
		us := service.NewUserService(uR)
		userController := controller.NewUserController(us)
		userRouterPublic := Router.Group("/users",userController.Register)
		{
			userRouterPublic.POST("/register")
			userRouterPublic.POST("/login")
			userRouterPublic.POST("/otp")
		}

		// private router
		userRouterPrivate := Router.Group("/users")
		//userRouterPrivate.Use(Permission())
		{
			userRouterPrivate.GET("/profile")
		}
		
	}
