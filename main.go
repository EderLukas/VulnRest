// @title           Swagger VulnAPI
// @version         1.0
// @description     This is the documentation to a vulnerable RestApi
// @termsOfService  http://swagger.io/terms/
// @license.name  MIT
// @license.url   https://opensource.org/licenses/MIT
// @host      localhost:8080
// @BasePath  /
// @schemes http
package main

import (
	"net/http"
	"log"
	"github.com/gin-gonic/gin"
 	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/files"
	_ "led/vulnapi/docs"
	_ "led/vulnapi/doctypes"
)

// @Summary      greets user
// @Description  returns greeting
// @Produce      json
// @Success      200  {object} doctypes.HelloEndPointResponse
// @Router       /hallo [get]
func HelloEndPoint(c *gin.Context) {
	log.Println("send Hello Response")
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello!",
	})
}

func main () {
	router := gin.Default()
    
	router.GET("/hello", HelloEndPoint)
	
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
    router.Run("localhost:8080")
}

