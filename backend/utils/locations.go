package utils

import (
	"fmt"
	"os"

	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type LocationResponse struct {
	Feature []struct {
		Geometry struct {
			Type        string `json:"Type"`
			Coordinates string `json:"Coordinates"`
		} `json:"Geometry"`
		Property struct {
			Address string `json:"Address"`
		} `json:"Property"`
	} `json:"Feature"`
}

func NearbyLocationsQuery(lat, lon float64) []qm.QueryMod {
	conditions := []qm.QueryMod{
		qm.Where("latitude >= ?", lat-0.0025),
		qm.Where("latitude <= ?", lat+0.0025),
		qm.Where("longitude >= ?", lon-0.0025),
		qm.Where("longitude <= ?", lon+0.0025),
		qm.Limit(20),
	}
	return conditions
}

func YahooApiUrl(lat, lon string) string {
	return fmt.Sprintf(
		"%s?lat=%s&lon=%s&appid=%s&output=json",
		os.Getenv("YAHOO_API_REVERSE_GEO_CODER_END_POINT"),
		lat,
		lon,
		os.Getenv("YAHOO_API_KEY"),
	)
}
