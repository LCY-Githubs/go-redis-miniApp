package router

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func init() {
	routerSlice = append(routerSlice, TestController)
}

func TestController(v1 *gin.RouterGroup) {
	r := v1.Group("/hello")
	{
		r.GET("", handle)
	}
}

func handle(c *gin.Context) {
	fmt.Println("hello world")
	c.JSON(http.StatusOK, "hello")
}
