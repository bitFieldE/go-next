package v1

import (
	"net/http"

	_ "github.com/lib/pq"

	"github.com/gin-gonic/gin"
)

func GetTop(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{"message": "Hello, Go!"})
}
