package middlewares

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func PrintURLMiddleware(c *gin.Context) {
	fmt.Println(c.Request.Method, c.Request.URL.Path)
	c.Next()
}
