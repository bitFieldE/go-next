package v1

import (
	"context"
	"encoding/json"
	"net/http"

	_ "github.com/lib/pq"

	db "github.com/bitFieldE/go-next/backend/db"
	models "github.com/bitFieldE/go-next/backend/models"
	utils "github.com/bitFieldE/go-next/backend/utils"
	"github.com/gin-gonic/gin"
)

var locationResponse utils.LocationResponse

func GetCurrentLocation(c *gin.Context) {
	url := utils.YahooApiUrl(c.Query("lat"), c.Query("lon"))

	db := db.ConnectDB()

	// MEMO:現在位置が誤差の範囲内にある地域（選択候補）を取得する
	locations, _ := models.Locations(
		utils.NearbyLocationsQuery(c.GetFloat64("lat"), c.GetFloat64("lon"))...,
	).All(context.Background(), db)

	defer db.Close()

	if locations != nil {
		c.IndentedJSON(http.StatusOK, locations)
		return
	}

	resp, err := http.Get(url)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, errorResponse(err))
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
