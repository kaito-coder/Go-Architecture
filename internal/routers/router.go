package routers

import (
	"fmt"

	"github.com/gin-gonic/gin"
	c "github.com/kaito-coder/go-ecommerce-architecture/internal/controller"
	"github.com/kaito-coder/go-ecommerce-architecture/internal/middlewares"
)
func AA() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		fmt.Println("Before --> AA")
		ctx.Next()
		fmt.Println("After --> AA")
	}
}
func BB() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		fmt.Println("Before --> BB")
		ctx.Next()
		fmt.Println("After --> BB")
	}
}
func CC(c *gin.Context) {
	fmt.Println("Before --> CC")
	c.Next()
	fmt.Println("After --> CC")
}


func NewRouter() *gin.Engine {
	r := gin.Default()
	r.Use(middlewares.AuthMiddleware(),AA(), BB(), CC)
	v1 := r.Group("/v1")
	{
		v1.GET("/user/:uid", c.NewUserController().GetUserById)
		v1.GET("/ping", pong)
		v1.GET("/ping/:uid", pong)
		v1.POST("/ping", pong)
		v1.PUT("/ping/:uid", pong)
		v1.DELETE("/ping/:uid", pong)
	}
	return r
}
func pong(c *gin.Context) {
	query := c.Query("name")
	defaultQuery := c.DefaultQuery("name", "default")
	param := c.Param("uid")
	c.JSON(200, gin.H{
		"message": "pong" + query + param + defaultQuery,
	})
}