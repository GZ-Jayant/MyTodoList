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
