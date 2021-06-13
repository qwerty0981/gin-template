package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// You can register routes like this
func init() {
	router.GET("/test", func(c *gin.Context) {
		c.String(http.StatusOK, "This is a test message")
	})
}
