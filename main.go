package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func init() {
	content, err := ioutil.ReadFile(".config")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(content))
	
}

func main() {

	/* 	r := gin.Default()
	   	r.LoadHTMLGlob("templates/*.tmpl")

	   	r.GET("/", func(c *gin.Context) {
	   		c.HTML(http.StatusOK, "index.tmpl", gin.H{})
	   	})

	   	r.POST("/gerarcodigo", func(c *gin.Context) {
	   		usuario := c.PostForm("usuario")
	   		enviar(usuario)
	   		c.String(http.StatusOK, "teste", usuario)
	   	})
	   	r.Run() */

}
