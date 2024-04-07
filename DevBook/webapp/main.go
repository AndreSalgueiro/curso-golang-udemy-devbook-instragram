package main

import (
	"fmt"
	"log"
	"net/http"
	"webapp/src/config"
	"webapp/src/coockies"
	"webapp/src/router"
	"webapp/src/utils"
)

// função criada para ser executada apenas uma vez para gerar a chave aleatoria do HASKEY e BLOCKKEY
/*func init() {
	hashKey := hex.EncodeToString(securecookie.GenerateRandomKey(16))
	fmt.Println(hashKey)

	blockKey := hex.EncodeToString(securecookie.GenerateRandomKey(16))
	fmt.Println(blockKey)
}*/

func main() {
	config.Carregar()
	cookies.Configurar()
	utils.CarregarTemplates() //carega todas as páginas de html na memoria para serem utilizadas
	r := router.Gerar()

	fmt.Printf("Escutando na porta %d\n", config.Porta)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Porta), r))
}
