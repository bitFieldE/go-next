package routers

import (
	"log"

	v1 "github.com/bitFieldE/go-next/backend/controller/api/v1"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func printReqCtxMiddleware(c *gin.Context) {
	log.Printf("Reaquest method: %s, path: %s", c.Request.Method, c.Request.URL.Path)
	c.Next()
}

func InitRouter(r *gin.Engine) {
	// CORSの設定
	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"GET", "POST", "DELETE", "OPTIONS"},
		AllowHeaders: []string{"*"},
	}))

	r.Use(printReqCtxMiddleware)
	namespace := r.Group("/api/v1")
	{
		namespace.GET("/", v1.GetTop)
		namespace.GET("/users", v1.GetUsers)
	}

}
