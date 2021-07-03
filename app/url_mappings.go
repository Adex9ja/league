package app

import "github.com/Adex9ja/solution/controllers"

func mapUrls() {
	router.POST("/echo", controllers.Echo)
	router.POST("/flatten", controllers.Flatten)
	router.POST("/sum", controllers.Sum)
	router.POST("/multiply", controllers.Multiply)
	router.POST("/invert", controllers.Invert)
}
