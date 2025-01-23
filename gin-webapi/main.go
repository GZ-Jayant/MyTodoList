package main

import (
	"gin-webapi/db"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	db.Connect()
	defer db.Close()

	r.Use(CORSMiddleware())

	api := r.Group("/api")
	{
		task := api.Group("/task")
		{
			task.POST("/create", func(c *gin.Context) {
				msg := "create task"
				tid, err := db.TaskCreate(c)
				if err != nil {
					msg = "create task error"
				}
				c.JSON(200, gin.H{
					"message": msg,
					"taskID":  tid,
				})
			})

			task.PUT("/update/:taskID", func(c *gin.Context) {
				msg := "update task"
				err := db.TaskUpdate(c)
				if err != nil {
					msg = "update task error"
				}
				c.JSON(200, gin.H{
					"message": msg,
				})
			})

			task.PUT("/delete/:taskID", func(c *gin.Context) {
				msg := "delete task"
				err := db.TaskDelete(c)
				if err != nil {
					msg = "delete task error"
				}
				c.JSON(200, gin.H{
					"message": msg,
				})
			})

			task.POST("/select", func(c *gin.Context) {
				var msg = "select task"

				taskList, err := db.TaskSelect(c)
				if err != nil {
					msg = "select task error"
				}

				c.JSON(200, gin.H{
					"message":  msg,
					"taskList": taskList,
				})
			})

		}

	}

	r.Run("192.168.50.118:8888")
}

/*
用于处理跨域资源共享（CORS）的中间件
CORS 是一种安全机制，允许服务器指定哪些域可以访问其资源

首先通过 c.Writer.Header().Set 方法设置多个 CORS 相关的 HTTP 头
Access-Control-Allow-Origin 头被设置为 "*", 允许任何域访问资源
Access-Control-Allow-Credentials 头被设置为 "true"，允许浏览器发送包含凭据的请求
Access-Control-Allow-Headers 头指定了允许的请求头列表，包括 Content-Type, Authorization 等
Access-Control-Allow-Methods 头指定了允许的 HTTP 方法，包括 POST, OPTIONS, GET, PUT

随后检查请求方法是否为 OPTIONS
OPTIONS 方法通常用于浏览器在实际请求之前发送的预检请求
如果请求方法是 OPTIONS，函数调用 c.AbortWithStatus(204)，返回一个 204 No Content 状态码，并终止请求处理
否则，函数会调用 c.Next()，将控制权交给下一个中间件或处理程序。
*/
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
