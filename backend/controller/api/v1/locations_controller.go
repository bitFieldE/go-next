package v1

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	_ "github.com/lib/pq"

	db "github.com/bitFieldE/go-next/backend/db"
	models "github.com/bitFieldE/go-next/backend/models"
	"github.com/gin-gonic/gin"
)

var locationResponse models.LocationResponse

func GetCurrentLocation(c *gin.Context) {
	url := fmt.Sprintf(
		"%s?lat=%s&lon=%s&appid=%s&output=json",
		os.Getenv("YAHOO_API_REVERSE_GEO_CODER_END_POINT"),
		c.Query("lat"),
		c.Query("lon"),
		os.Getenv("YAHOO_API_KEY"),
	)

	db := db.ConnectDB()
	locations, _ := models.Locations().All(context.Background(), db)
	defer db.Close()

	if locations != nil {
		c.IndentedJSON(http.StatusOK, gin.H{"location": url})
		return
	}

	resp, err := http.Get(url)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	defer resp.Body.Close()

	if err = json.NewDecoder(resp.Body).Decode(&locationResponse); err != nil {
		c.IndentedJSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	c.IndentedJSON(http.StatusOK, locationResponse)
}

func ResiterLocation(c *gin.Context) {
	db := db.ConnectDB()
	defer db.Close()
}
