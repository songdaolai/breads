package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
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
					put: "新闻标题11199999999999999，这是我修改的东西",
					pht: "新闻标题222",
				},
			},
		})
	})
	r.Run()
}
