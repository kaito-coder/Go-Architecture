package admin

import "github.com/gin-gonic/gin"

type UserRouter struct {}

func (ur *UserRouter) InitUserRouter(Router *gin.RouterGroup) {

	// private router
	userRouterPrivate := Router.Group("/admin/users")
	//userRouterPrivate.Use(Limiter())
	//userRouterPrivate.Use(Permission())
	//userRouterPrivate.Use(Permission())
	{
		userRouterPrivate.GET("/active_user")
	}
	
}
