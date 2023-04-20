package main

import (
	"github.com/RkAirforce/go-next/backend/config"
	db "github.com/RkAirforce/go-next/backend/db"
	routers "github.com/RkAirforce/go-next/backend/router"

	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	defer db.Close()
	router := gin.Default()
	routers.InitRouter(router)
	router.Run(":" + config.Server.Port)
}
