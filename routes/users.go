package routes

import ("github.com/gin-gonic/gin"
        "gin-gorm-rest-docker/controller"
)

func GetUsers(router *gin.Engine){
	router.GET("/", controller.UserController)
}