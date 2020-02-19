package controller

import (
	"github.com/gin-gonic/gin"
	"model"
	"net/http"
)

func Login(c *gin.Context) {
	if err := model.LoadData(c); err == nil {
		pass := model.GetPassword()
		err := model.UserInfo.GetUserByUsername()
		if err != nil {
			Response(c, http.StatusOK, -1, "用户不存在", err.Error())
			return
		}
		if model.CheckPassword(pass) {
			model.SetAuthCode()
			Response(c, http.StatusOK, http.StatusOK, "登录成功", model.UserInfo)
			return
		} else {
			Response(c, http.StatusOK, -2, "密码错误", err.Error())
			return
		}
	} else {
		Response(c, http.StatusBadGateway, http.StatusBadGateway, "参数缺失", err.Error())
		return
	}
}
