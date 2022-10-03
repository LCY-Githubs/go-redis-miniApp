package controller

import (
	"fmt"
	"net/http"
	"src/app/apis"

	"src/app/util"

	"github.com/gin-gonic/gin"
)

type InterviewController struct {
	apis.Api
}

func (c InterviewController) Get(name string) string {
	return "Hello World"
}

func (c InterviewController) GetName(r *gin.Context) {
	fmt.Println("Hello World2")
	r.JSON(http.StatusOK, "hello")
}

func (c InterviewController) GetId(r *gin.Context) {
	r.JSON(http.StatusOK, "hello getId")
}

func (c InterviewController) HandleMsg(r *gin.Context) {
	fmt.Println("Hello World handleMsg")
	// r.JSON(http.StatusOK, "hello")
	util.OK(r, "Hello", "success")
	// res.SetCode(200)
	// res := util.Response{200, "Hello", "success"}
	// r.JSON(http.StatusOK, res)
	// c.OK(res, "查询成功")
}
