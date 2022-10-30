package util

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func OK(r *gin.Context, msg string, data interface{}) {
	r.JSON(http.StatusOK, CommonRes{Msg: msg, Data: data})
}

func FAIL(r *gin.Context) {
	r.JSON(http.StatusInternalServerError, CommonRes{Msg: "FAIL"})
}
