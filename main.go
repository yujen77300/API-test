package main

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/gin-gonic/gin/binding"
	"github.com/yujen77300/API-test/middleware"
	"github.com/yujen77300/API-test/database"
	"github.com/yujen77300/API-test/database/migration"
	"github.com/yujen77300/API-test/router"
)

func main() {
	database.DatabaseInit()
	database.RedisInit()
	migration.Migration()

	r := gin.Default()

	if v, ok := binding.Validator.Engine().(*validator.Validate);ok{
		v.RegisterValidation("pwdvaldation",middleware.PwdValidation)
	}

	v1 := r.Group("/v1")
	router.RouterInit(v1)

	r.Run(":8080")
}
