package main

import (
	db "github.com/RkAirforce/go-next/backend/postgres/db"
	routers "github.com/RkAirforce/go-next/backend/router"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	db.InitDB()
	routers.InitRouter(router)
	router.Run(":3000")
}
