package util

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var Default = &response{}

//通常成功数据处理
func OK(c *gin.Context, data interface{}, msg string) {
	r := Default.Clone()
	r.SetData(data)
	r.SetSuccess(true)
	if msg != "" {
		r.SetMsg(msg)
	}
	r.SetCode(http.StatusOK)
	c.Set("result", r)
	c.Set("status", http.StatusOK)
	c.AbortWithStatusJSON(http.StatusOK, r)
}
