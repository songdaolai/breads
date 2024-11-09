package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()

	r.LoadHTMLGlob("templates/**/*")
	r.GET("/user", func(c *gin.Context) {
		c.HTML(http.StatusOK, "default/user.html", gin.H{})
	})

	r.POST("/doAddUser", func(c *gin.Context) {
		username := c.PostForm("username")
		password := c.PostForm("password")
		c.JSON(http.StatusOK, gin.H{
			"username": username,
			"password": password,
		})
	})
	r.Run()
}
