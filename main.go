package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*.tmpl")

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{})
	})

	r.POST("/gerarcodigo", func(c *gin.Context) {
		usuario := c.PostForm("usuario")
		c.String(http.StatusOK, "teste", usuario)
	})
	r.Run()
}
