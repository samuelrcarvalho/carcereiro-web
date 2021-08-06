package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func grant(tabelas []string, usuario string) {
	// Connection target
	dsg := configs["target_database_user"] + ":" + configs["target_database_pwd"] + "@tcp(" + configs["target_database_host"] + ":" + configs["target_database_port"] + ")/" + configs["target_database_schema"] + "?charset=utf8mb4&parseTime=True&loc=Local"
	dbtarget, err := gorm.Open(mysql.Open(dsg), &gorm.Config{})
	if err != nil {
		panic("failed to connect target database")
	}
	for _, item := range tabelas {
		dbtarget.Exec("GRANT SELECT ON TABLE `" + configs["target_database_schema"] + "`.`" + item + "` TO '" + usuario + "'@'%'")
	}
}

//Interface para permitir alteracao de table fora do plural do Gorm
type Tabler interface {
	TableName() string
}

//Method para permitir alteracao de table fora do plural do Gorm
func (User) TableName() string {
	return "user"
}

type User struct {
	User string
}

func usuarioExiste(recebeUsuario string) bool {
	var usuario []User
	existe := false
	dscu := configs["target_database_user"] + ":" + configs["target_database_pwd"] + "@tcp(" + configs["target_database_host"] + ":" + configs["target_database_port"] + ")/mysql?charset=utf8mb4&parseTime=True&loc=Local"
	dbcheckuser, err := gorm.Open(mysql.Open(dscu), &gorm.Config{})
	if err != nil {
		panic("failed to connect target database")
	}
	result := dbcheckuser.Where("user = ?", recebeUsuario).First(&usuario)
	if result.RowsAffected != 0 {
		existe = true
	}
	return existe
}

type Importanttable struct {
	Nome string `gorm:"primaryKey"`
}

// tabelaRestrita autoriza liberar acesso em tabelas que não tem restrição
func tabelaRestrita(db *gorm.DB, tabelasSolicitacao []string) bool {
	var tabelaImportante []Importanttable
	temrestricao := true
	for _, item := range tabelasSolicitacao {
		result := db.Where("nome = ?", item).First(&tabelaImportante)
		if result.RowsAffected == 0 {
			temrestricao = false
		} else {
			temrestricao = true
			break
		}
	}
	return temrestricao
}
