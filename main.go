package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type act struct {
	put string
	pht string
}

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/**/*")
	r.GET("admin/news", func(c *gin.Context) {
		c.HTML(http.StatusOK, "news.html", gin.H{
			"title": "首页",
			"newt": []interface{}{
				&act{
					put: "新闻标题111",
					pht: "新闻标题222",
				},
			},
		})
	})
	r.Run()
}
