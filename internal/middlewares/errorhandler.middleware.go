package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/kaito-coder/go-ecommerce-architecture/pkg/response"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token == "" {
			response.FailResponse(c, response.ErrAuthFail, "Unauthorized")
			c.Abort()
			return
		}
		c.Next()	
	}
}
