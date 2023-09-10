package v1

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	_ "github.com/lib/pq"

	"github.com/gin-gonic/gin"
)

func SignUpHandler(c *gin.Context) {
	tokenLifeTime, err := strconv.Atoi(os.Getenv("TOKEN_LIFETIME"))
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"err": err})
	}
	fmt.Println(time.Now().Add(time.Hour * time.Duration(tokenLifeTime)).Unix())
	c.IndentedJSON(http.StatusBadRequest, gin.H{"token": time.Now().Add(time.Hour * time.Duration(tokenLifeTime)).Unix()})
}

func LoginHandler(context *gin.Context) {

}
