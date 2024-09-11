package initialize

import (
	"github.com/gin-gonic/gin"
	"github.com/kaito-coder/go-ecommerce-architecture/global"
	"github.com/kaito-coder/go-ecommerce-architecture/internal/routers"
)



func InitRouter() *gin.Engine {
	var r *gin.Engine
	if global.Config.Server.Mode == "dev"{
		gin.SetMode(gin.DebugMode)
		gin.ForceConsoleColor()
		r = gin.Default()
	}else {
		gin.SetMode(gin.ReleaseMode)
		r = gin.New()
	}
	// middlewares
	//r.Use(middleware.limiter())
	//r.Use(middleware.Cors())
	//r.Use(middleware.Logger())
	//r.Use(middleware.Recovery())
	adminRouter := routers.RouterGroupApp.Admin
	userRouter := routers.RouterGroupApp.User

	MainGroup := r.Group("/v1")
	{
		MainGroup.GET("/checkStatus")
	}
	{
		userRouter.InitUserRouter(MainGroup)
	}
	{
		adminRouter.InitAdminRouter(MainGroup)
		adminRouter.InitUserRouter(MainGroup)
	}
	return r
}

