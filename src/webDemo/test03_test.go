package webDemo

import (
	_ "net/http"
	"testing"

	"github.com/gin-gonic/gin"
)

// func Test03(t *testing.T) {
// 	r := gin.Default()
// 	r.GET("/ping", func(c *gin.Context) {
// 		c.JSON(http.StatusOK, gin.H{
// 			"message": "pong",
// 		})
// 	})
// 	r.Run()
// }

func Test04(t *testing.T) {
	r := gin.Default()
	r.StaticFile("logo.jpg", "E:/photo/有用的图片/8KRAWPIC0000084.jpg")
	r.Run()
}
