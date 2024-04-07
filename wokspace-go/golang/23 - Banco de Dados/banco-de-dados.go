package main

import (
	"database/sql"
	"fmt"
	"log"

	/*usando o pacote de forma implícita isso ocorre quando precisa usar de forma indiretamente,
	 quem vai usar este drive é o pacote database/sql*/
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	stringConexao := "golang:2810@/dev_book?charset=utf8&parseTime=True&loc=Local" //usuario:senha@/nome banco
	db, erro := sql.Open("mysql", stringConexao)
	if erro != nil {
		log.Fatal(erro)
	}
	defer db.Close()//fechando a conexão do banco depois que acabar de ulilizar
	
	if erro = db.Ping(); erro != nil {
		log.Fatal(erro)
	}

	fmt.Println("Conexão está aberta")

	linhas, erro := db.Query("select * from usuarios")
	if erro != nil {
		log.Fatal(erro)
	}
	// as linhas retornadas deuma consulta sql, é um cursor que precisa ser fechada
	defer linhas.Close()

	fmt.Println(linhas)
}
