package router

import (
	"fmt"
	"src/app/controller"

	"github.com/gin-gonic/gin"
)

func init() {
	routerSlice = append(routerSlice, InterviewApiController)
}

//建立外部访问系统的接口
func InterviewApiController(v1 *gin.RouterGroup) {
	api := controller.InterviewController{}
	r := v1.Group("/interview")
	fmt.Printf("interview")
	{
		r.GET("", api.GetName)
		r.GET("/:id", api.GetId)
		r.GET("/handle", api.HandleMsg)
	}

}
