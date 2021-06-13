package routes

import "github.com/gin-gonic/gin"

func initRouter() *gin.Engine {
	// Here you can customize the server-wide router middleware
	r := gin.Default()

	return r
}

var router = initRouter()

func GetRouter() *gin.Engine {
	return router
}
