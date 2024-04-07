package banco

import (
	"api/src/config"
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func Conectar() (*sql.DB, error) {
	
	db, erro := sql.Open("mysql", config.StringConexaoBD)
	if erro != nil {
		return nil, erro
	}

	//verificando se a conexão com o banco está aberta utilizando a função ping()
	if erro = db.Ping(); erro != nil {
		db.Close()
		return nil, erro
	}
	return db, nil
}