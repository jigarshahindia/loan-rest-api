package middelware

import (
	"github.com/gin-gonic/gin"
)

/*
Already exist
*/
func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader != "Basic dGVzdDp0ZXN0" {
			c.JSON(401, "not authorized")
		} else {
			c.Next()
		}

	}
}
