package main

import (
	"github.com/gin-gonic/gin"
	"surround/database"
	"surround/router"
)

func main() {
	defer database.Database.Close()
	defer database.Cache.Close()

	r := router.InitRouter()
	gin.SetMode(gin.DebugMode)
	err := r.Run(":5000")

	if err != nil {
		panic(err)
	}
}
