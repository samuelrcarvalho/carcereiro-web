package main

//Interface para permitir alteracao de table fora do plural do Gorm
type Tabler interface {
	TableName() string
}

//Method para permitir alteracao de table fora do plural do Gorm
func (tables_priv) TableName() string {
	return "tables_priv"
}

type tables_priv struct {
	User       string
	Table_name string
	Host       string
	Db         string
}

func revokeUserPrivileges() {
	listaUsuarios := listaDeUsuariosParaRevogar()
	var tabelasRevoke []tables_priv
	for _, item := range listaUsuarios {
		_ = dbmysql.Where("user = ?", item).Find(&tabelasRevoke)
		for _, item2 := range tabelasRevoke {
			dbmysql.Exec("REVOKE ALL ON " + item2.Db + "." + item2.Table_name + " FROM '" + item2.User + "'@'" + item2.Host + "';")
		}
	}
}

type usersExcept struct {
	User string
}

func listaDeUsuariosParaRevogar() []string {
	var listao []string
	var pegaLista []tables_priv
	var userRestritos []usersExcept
	dbmysql.Distinct("User").Find(&pegaLista)
	for _, item := range pegaLista {
		result := db.Where("user = ?", item.User).First(&userRestritos)
		if result.RowsAffected != 1 {
			listao = append(listao, item.User)
		}
	}
	return listao
}
