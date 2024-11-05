package main

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/gin-gonic/gin"
)

type act struct {
	Put string
	Pht string
}

func Prithen(str1 string, str2 string) string {
	fmt.Println(str1, str2)
	return str1 + "---" + str2
}
func main() {
	r := gin.Default()
	r.SetFuncMap(template.FuncMap{
		"Prithee": Prithen,
	})
	r.LoadHTMLGlob("templates/**/*")
	r.GET("admin/news", func(c *gin.Context) {
		c.HTML(http.StatusOK, "news.html", gin.H{
			"title": "首页",
			"but":   "新闻",
			"news": &act{
				Put: "新闻标题222",
				Pht: "新闻标题333",
			},
		})
	})

	r.Run()
}
