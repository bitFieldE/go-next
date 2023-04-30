package main

import (
	"github.com/bitFieldE/go-next/backend/config"
	routers "github.com/bitFieldE/go-next/backend/router"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	routers.InitRouter(router)
	router.Run(":" + config.Server.Port)
}
