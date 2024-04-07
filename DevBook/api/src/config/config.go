package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	//String de conexão com banco de dados
	StringConexaoBD = ""
	//Porta onde a API estar rodando
	Porta = 0

	//Secretkey é a chave que vai ser usada para assinar o token
	SecretKey []byte
)

// Carregar carrega as variaveis de ambiente da API
func Carregar() {
	var erro error

	if erro = godotenv.Load(); erro != nil {
		log.Fatal(erro)
	}
	//extrai a porta do arqiovo .env e transforma para inteiro
	Porta, erro = strconv.Atoi(os.Getenv("API_PORT"))
	//se não tiver uma porta declarada no arquivo de cong .env então pega a porta padrão 9000
	if erro != nil {
		Porta = 9000
	}
	//usuario:senha@/nome banco
	StringConexaoBD = fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True&loc=Local",
	os.Getenv("DB_USUARIO"),
	os.Getenv("DB_SENHA"),
	os.Getenv("DB_NOME"))

	SecretKey = []byte(os.Getenv("SECRET_KEY"))
}
