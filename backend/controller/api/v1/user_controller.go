package v1

import (
	"context"
	"net/http"

	_ "github.com/lib/pq"

	db "github.com/bitFieldE/go-next/backend/db"
	models "github.com/bitFieldE/go-next/backend/models"
	"github.com/gin-gonic/gin"
)

func GetUsers(c *gin.Context) {
	db := db.ConnectDB()
	users, _ := models.Users().All(context.Background(), db)
	c.IndentedJSON(http.StatusOK, gin.H{"users": users})
}
