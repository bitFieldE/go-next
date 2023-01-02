package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func top(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{"message": "pong"})
}

func main() {
	router := gin.Default()
	router.GET("/", top)

	router.Run(":3000")
}
