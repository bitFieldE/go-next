package v1

import (
	"context"
	"net/http"
	"strconv"

	_ "github.com/lib/pq"

	db "github.com/bitFieldE/go-next/backend/db"
	models "github.com/bitFieldE/go-next/backend/models"
	"github.com/gin-gonic/gin"
)

func GetUsers(c *gin.Context) {
	db := db.ConnectDB()
	users, _ := models.Users().All(context.Background(), db)
	defer db.Close()
	c.IndentedJSON(http.StatusOK, gin.H{"users": users})
}

func GetUser(c *gin.Context) {
	db := db.ConnectDB()
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, errorResponse(err))
		return
	}
	user, _ := models.FindUser(context.Background(), db, id)
	defer db.Close()
	c.IndentedJSON(http.StatusOK, gin.H{"user": user})
}
