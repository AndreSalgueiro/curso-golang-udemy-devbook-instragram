package banco

import ("database/sql"
	/*usando o pacote de forma implícita isso ocorre quando precisa usar de forma indiretamente,
	 quem vai usar este drive é o pacote database/sql*/
	 _ "github.com/go-sql-driver/mysql"
)
func Conectar() (*sql.DB, error) {
	stringConexao := "golang:2810@/dev_book?charset=utf8&parseTime=True&loc=Local" //usuario:senha@/nome banco

	db, erro := sql.Open("mysql", stringConexao)
	if erro != nil {
		return nil, erro
	}

	if erro = db.Ping(); erro != nil {
		return nil, erro
	}

	return db, nil
}