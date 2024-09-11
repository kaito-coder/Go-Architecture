package admin

import "github.com/gin-gonic/gin"

type AdminRouter struct {}

func (ur *AdminRouter) InitAdminRouter(Router *gin.RouterGroup) {

	adminRouterPublic := Router.Group("/admin")
	{
		adminRouterPublic.POST("/login")
	}
	// private router
	// adminRouterPrivate := Router.Group("/admin/users")
	// //userRouterPrivate.Use(Limiter())
	// //userRouterPrivate.Use(Permission())
	// //userRouterPrivate.Use(Permission())
	// {
	// 	adminRouterPrivate.GET("/active_user")
	// }
	
}
