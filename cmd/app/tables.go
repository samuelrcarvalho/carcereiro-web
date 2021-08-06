package main

import (
	"gorm.io/gorm"
)

type Table struct {
	Nome string `gorm:"primaryKey"`
}

func tabelaTabelas(db *gorm.DB) []string {
	var listao []string
	var pegaLista []Table
	db.Find(&pegaLista)

	for _, item := range pegaLista {
		listao = append(listao, item.Nome)
	}
	return listao

}
