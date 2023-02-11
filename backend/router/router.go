package routers

import (
	"log"

	v1 "github.com/RkAirforce/go-next/backend/controller/api/v1"
	"github.com/gin-gonic/gin"
)

func printReqCtxMiddleware(c *gin.Context) {
	log.Printf("Reaquest method: %s, path: %s", c.Request.Method, c.Request.URL.Path)
	c.Next()
}

func InitRouter(r *gin.Engine) {
	r.Use(printReqCtxMiddleware)
	namespace := r.Group("/api/v1")
	{
		namespace.GET("/", v1.GetTop)
	}

}
