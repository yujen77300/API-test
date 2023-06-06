package router

import (
	"github.com/gin-gonic/gin"
	"github.com/yujen77300/API-test/service"
)

func RouterInit(r *gin.RouterGroup) {
	user := r.Group("/user")

	user.GET("/all", service.GetAllUsers)
	user.GET("/:id", service.GetUserById)
	user.DELETE("/:id", service.DeleteUser)
}

func RouterUpdate(r *gin.RouterGroup) {
	user := r.Group("/user")

	user.POST("/", service.CreateUser)
	user.POST("/verify", service.VerifyUser)
}
