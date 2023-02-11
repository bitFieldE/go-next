package main

import (
	routers "github.com/RkAirforce/go-next/backend/router"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	routers.InitRouter(router)
	router.Run(":3000")
}
