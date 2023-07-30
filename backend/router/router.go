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
		auth := namespace.Group("/auth")
		{
			auth.GET("/signup", v1.SignUpHandler)
			auth.POST("/login", v1.LoginHandler)
		}
		top := namespace.Group("/top")
		{
			top.GET("/", v1.GetTop)
		}
		users := namespace.Group("/users")
		{
			users.GET("", v1.GetUsers)
			users.GET("/:id", v1.GetUser)
		}
		location := namespace.Group("/location")
		{
			location.GET("/currentLocation", v1.GetCurrentLocation)
		}
	}

}
