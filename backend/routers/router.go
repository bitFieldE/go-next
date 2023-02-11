package routers

import (
	"log"

	"github.com/gin-gonic/gin"
)

func printReqCtxMiddleware(c *gin.Context) {
	log.Printf("Reaquest method: %s, path: %s", c.Request.Method, c.Request.URL.Path)
	c.Next()
}

func InitRouter(r *gin.Engine) {
	r.Use(printReqCtxMiddleware)

	namespace := r.Group("/api/v1")

}
