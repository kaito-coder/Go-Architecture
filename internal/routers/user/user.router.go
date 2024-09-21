package user

import (
	"github.com/gin-gonic/gin"
	"github.com/kaito-coder/go-ecommerce-architecture/internal/wire"
)

	type UserRouter struct {}

	func (ur *UserRouter) InitUserRouter(Router *gin.RouterGroup) {
		// public router 
		// non dependency injection
		// uR := repo.NewUserRepository()
		// us := service.NewUserService(uR)
		// userController := controller.NewUserController(us)
		userController, _ := wire.InitUserRouterHandler()
		userRouterPublic := Router.Group("/users")
		{
			userRouterPublic.POST("/register", userController.Register)
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
