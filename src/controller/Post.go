package controller

import (
	"github.com/gin-gonic/gin"
	"model"
	"net/http"
	"strconv"
)

func List(c *gin.Context) {
	err := model.GetPostList()
	var msg string
	if err != nil {
		msg = err.Error()
	} else {
		msg = "获取成功"
	}
	Response(c, http.StatusOK, http.StatusOK, msg, model.PostInfo)
}
func Get(c *gin.Context) {
	id, _ := c.GetQuery("id")
	pid, err := strconv.Atoi(id)
	if err != nil {
		Response(c, http.StatusBadGateway, http.StatusBadGateway, "参数缺失", err.Error())
		return
	}
	var post model.Post
	post, _ = model.GetDetail(pid)
	Response(c, http.StatusOK, http.StatusOK, "获取成功", post)
}
func Add(c *gin.Context) {
	post := model.Post{}
	_ = c.ShouldBindJSON(&post)
	err := post.Add()
	if err != nil {
		Response(c, http.StatusOK, http.StatusOK, "新增失败", err.Error())
		return
	}
	Response(c, http.StatusOK, http.StatusOK, "新增成功", nil)
}
func Update(c *gin.Context) {
	post := model.Post{}
	_ = c.ShouldBindJSON(&post)
	res, err := post.Update()
	if err != nil || res == 0 {
		Response(c, http.StatusOK, http.StatusOK, "更新失败", nil)
		return
	}
	Response(c, http.StatusOK, http.StatusOK, "更新成功", nil)
}
func Delete(c *gin.Context) {
	id, _ := c.GetQuery("id")
	pid, err := strconv.Atoi(id)
	if err != nil {
		Response(c, http.StatusBadGateway, http.StatusBadGateway, "参数缺失", err.Error())
		return
	}
	res, err := model.Delete(pid)
	if err != nil || res == 0 {
		Response(c, http.StatusOK, http.StatusOK, "删除失败", nil)
		return
	}
	Response(c, http.StatusOK, http.StatusOK, "删除成功", nil)
}
