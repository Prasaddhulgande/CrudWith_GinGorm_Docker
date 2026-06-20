package main

import (
	"github.com/gin-gonic/gin"
	"gin-gorm-rest-docker/routes"
)
func main() {
    
	router := gin.Default()

	routes.UserRoute(router)

	router.Run(":8000")
}