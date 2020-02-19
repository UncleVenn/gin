package main

import (
	"controller"
	"github.com/gin-gonic/gin"
	"model"
	"net/http"
)

func main() {
	router := gin.Default()
	router.StaticFS("/static", http.Dir("static"))
	router.POST("/login", controller.Login)
	router.Use(Authorize())
	router.GET("/list", controller.List)
	router.GET("/get", controller.Get)
	router.POST("/add", controller.Add)
	router.POST("/update", controller.Update)
	router.GET("/delete", controller.Delete)
	_ = router.Run(":8080")
}
func Authorize() gin.HandlerFunc {
	return func(c *gin.Context) {
		model.UserInfo.AuthCode = c.GetHeader("AuthCode")
		if err := model.UserInfo.GetUserByAuthCode(); err != nil {
			controller.Response(c, http.StatusForbidden, http.StatusForbidden, "请先登录", err.Error())
			return
		}
	}
}
