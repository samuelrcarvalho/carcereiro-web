package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/spf13/viper"
)

func init() {
	// Reading configfile
	viper.SetConfigType("toml")
	viper.AddConfigPath(".")
	viper.SetConfigName(".env")
	viper.ReadInConfig()
}

var configs map[string]string

func main() {
	dsn := viper.GetString("user") + ":" + viper.GetString("password") + "@tcp(" + viper.GetString("host") + ":" + viper.GetString("port") + ")/carcereiro?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	configs = configDB(db)

	// starting Gin
	r := gin.Default()
	r.LoadHTMLGlob("templates/*.tmpl")

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{})
	})

	r.POST("/gerarcodigo", func(c *gin.Context) {
		usuario := c.PostForm("usuario")
		enviarCode(usuario, db)
		c.String(http.StatusOK, "", configs["rocket_url"])

	})
	r.Run()

}

type Config struct {
	Config string `gorm:"primaryKey"`
	Value  string
}

//map[string]string
func configDB(db *gorm.DB) map[string]string {
	configs := make(map[string]string)
	var listaTudo []Config
	db.Find(&listaTudo)
	for _, item := range listaTudo {
		configs[item.Config] = item.Value
	}
	return configs
}
