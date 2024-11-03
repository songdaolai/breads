package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type act struct {
	Put string
	Pht string
}

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/**/*")
	r.GET("admin/news", func(c *gin.Context) {
		c.HTML(http.StatusOK, "news.html", gin.H{
			"title": "首页",
			"newt": []interface{}{
				&act{
					Put: "新闻标题111",
					Pht: "新闻标题222",
				},
			},
		})
	})
	r.Run()
}
