package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default() //返回默认的路由信息

	// GET：请求方式；/book：请求的路径
	r.GET("/book", func(c *gin.Context) {
		// c.JSON：返回JSON格式的数据
		c.JSON(http.StatusOK, gin.H{
			"message": "GET book",
		})
	})
	r.PUT("/book", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"method": "PUT",
		})
	})
	r.POST("/book", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"method": "POST",
		})
	})
	r.DELETE("/book", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"method": "DELETE",
		})
	})
	//启动服务
	r.Run()
}
