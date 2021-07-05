//This takes care of how the application works and we can start the application
package app

import "github.com/gin-gonic/gin"

var (
	router *gin.Engine
)

func init() {
	router = gin.Default()
}

func StartApp() {
	mapUrls()
	if err := router.Run(":8080"); err != nil {
		panic(err)
	}
}
