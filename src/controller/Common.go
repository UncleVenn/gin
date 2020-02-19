package controller

import "github.com/gin-gonic/gin"

func Response(c *gin.Context, httpStatusCode int, code int, msg string, data interface{}) {
	c.AbortWithStatusJSON(httpStatusCode, gin.H{
		"code": code,
		"msg":  msg,
		"data": data,
	})
}
