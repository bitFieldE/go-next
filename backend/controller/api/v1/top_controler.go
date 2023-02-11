package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetTop(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{"message": "Hello, Gopher!"})
}
