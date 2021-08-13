package main

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/spf13/viper"
)

func init() {
	// Reading configfile
	viper.SetConfigType("toml")
	viper.AddConfigPath(os.Getenv("CARCEREIRO_HOME"))
	viper.SetConfigName(".env")
	viper.ReadInConfig()
}

var configs map[string]string
var usuario string
var lista []string

func main() {
	dsn := viper.GetString("user") + ":" + viper.GetString("password") + "@tcp(" + viper.GetString("host") + ":" + viper.GetString("port") + ")/carcereiro?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	configs = configDB(db)

	// starting Gin
	r := gin.Default()
	r.LoadHTMLGlob(os.Getenv("CARCEREIRO_HOME") + "/templates/*.tmpl")

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{})
	})

	r.POST("/gerarcodigo", func(c *gin.Context) {
		usuario := c.PostForm("usuario")
		enviarCode(usuario, db)
		c.HTML(http.StatusOK, "gerarcodigo.tmpl", gin.H{})
	})
	r.POST("/validacodigo", func(c *gin.Context) {
		codigo := c.PostForm("code")
		usuario := c.PostForm("user")
		if accessValid(db, usuario, codigo) == true {
			listao := tabelaTabelas(db)
			c.HTML(http.StatusOK, "tabelas.tmpl", gin.H{
				"entrada": listao,
			})
		} else {
			c.HTML(http.StatusOK, "error.tmpl", gin.H{
				"erro": "CÓDIGO ERRADO OU EXPIRADO",
			})
		}
	})
	r.POST("/libera", func(c *gin.Context) {
		var novoTicket string
		codigo := c.PostForm("code")
		usuario = c.PostForm("user")
		lista = c.PostFormArray("lista")
		justif := c.PostFormArray("justificativa")
		if accessValid(db, usuario, codigo) == true {
			if usuarioExiste(usuario) == true {
				if tabelaRestrita(db, lista) == true {
					novoTicket = abrirTicket(justif[0], false)
					c.HTML(http.StatusOK, "libera.tmpl", gin.H{
						"message": "Ticket " + novoTicket + " aberto no ServiceDesk e aguardando aprovação.",
					})
				} else {
					grant(lista, usuario)
					novoTicket = abrirTicket(justif[0], true)
					c.HTML(http.StatusOK, "libera.tmpl", gin.H{
						"message": "Ticket " + novoTicket + " criado no ServiceDesk para registro.",
					})
				}
			} else {
				c.HTML(http.StatusOK, "error.tmpl", gin.H{
					"erro": "USUARIO NÃO PERMITIDO PARA LIBERAÇÃO DE ACESSO",
				})
			}
		} else {
			c.HTML(http.StatusOK, "error.tmpl", gin.H{
				"erro": "CÓDIGO ERRADO OU EXPIRADO",
			})
		}
	})

	r.Run()
}

type Config struct {
	Config string `gorm:"primaryKey"`
	Value  string
}

func configDB(db *gorm.DB) map[string]string {
	configs := make(map[string]string)
	var listaTudo []Config
	db.Find(&listaTudo)
	for _, item := range listaTudo {
		configs[item.Config] = item.Value
	}
	return configs
}
