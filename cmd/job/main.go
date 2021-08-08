package main

import (
	"log"
	"os"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var db *gorm.DB
var dbmysql *gorm.DB

func init() {
	// Reading configfile
	viper.SetConfigType("toml")
	viper.AddConfigPath(".")
	viper.SetConfigName(".env")
	viper.ReadInConfig()

}

func main() {
	var err error
	var err2 error
	// Ocultar log ao não encontra em query
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			IgnoreRecordNotFoundError: true,
		},
	)
	// Conexão com a tabela carcereiro
	dsn := viper.GetString("user") + ":" + viper.GetString("password") + "@tcp(" + viper.GetString("host") + ":" + viper.GetString("port") + ")/carcereiro?charset=utf8mb4&parseTime=True&loc=Local"
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		panic("failed to connect database")
	}
	// Conexão com a tabela mysql
	dsmysql := viper.GetString("user") + ":" + viper.GetString("password") + "@tcp(" + viper.GetString("host") + ":" + viper.GetString("port") + ")/mysql?charset=utf8mb4&parseTime=True&loc=Local"
	dbmysql, err2 = gorm.Open(mysql.Open(dsmysql), &gorm.Config{})
	if err2 != nil {
		panic("failed to connect database")
	}
	_ = configDB(db)

	revokeUserPrivileges()
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
