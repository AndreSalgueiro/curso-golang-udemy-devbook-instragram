package main

import (
	"api/src/config"
	"api/src/router"
	"fmt"
	"log"
	"net/http"
)

/*func init() {
	//criando uma sicret em base 64
	//ESSE CÓDIGO NÃO É MAIS USADO, FOI SÓ PARA GERAR UMA CHAVE
	chave := make([]byte, 64)
	if _, erro := rand.Read(chave); erro != nil {
		log.Fatal(erro)
	}

	stringBase64 := base64.StdEncoding.EncodeToString(chave)
	fmt.Println(stringBase64)
}*/

func main() {
	//carregando as variáveis de ambiente
	config.Carregar()
	
	//Configurar as rotas da API
	r := router.Gerar()
	//sobe um servidor para a nossa API
	fmt.Printf("Escutando na porta %d", config.Porta)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Porta), r))
}