package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"

	"github.com/spf13/viper"
)

func init() {
	// Reading configfile
	viper.SetConfigType("toml")
	viper.AddConfigPath(".")
	viper.SetConfigName(".env")
	viper.ReadInConfig()
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

	db, err := sql.Open("mysql", viper.GetString("user")+":"+viper.GetString("password")+"@tcp("+viper.GetString("host")+":"+viper.GetString("port")+")/carcereiro")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	configs := configDB(db)
}

type Config struct {
	config string
	value  string
}

func configDB(db *sql.DB) map[string]string {
	rows, err := db.Query("SELECT * FROM config")
	if err != nil {
		log.Fatal(err)
	}
	configs := make(map[string]string)
	var config Config

	for rows.Next() {
		err := rows.Scan(&config.config, &config.value)
		if err != nil {
			log.Fatal(err)
		}
		configs[config.config] = config.value
	}
	return configs
}
