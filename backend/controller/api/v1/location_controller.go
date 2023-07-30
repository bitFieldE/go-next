package v1

import (
	"context"
	"net/http"

	_ "github.com/lib/pq"

	db "github.com/bitFieldE/go-next/backend/db"
	models "github.com/bitFieldE/go-next/backend/models"
	"github.com/gin-gonic/gin"
)

func GetCurrentLocation(c *gin.Context) {
	db := db.ConnectDB()
	location, _ := models.Locations().All(context.Background(), db)
	defer db.Close()
	c.IndentedJSON(http.StatusOK, gin.H{"location": location})
}
