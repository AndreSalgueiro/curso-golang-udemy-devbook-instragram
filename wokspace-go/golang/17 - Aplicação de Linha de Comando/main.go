package main

import (
	"fmt"
	"linha-de-comando/app"
	"os"
	"log"
)

func main() {
	fmt.Println("Ponto de partida")

	aplicacao := app.Gerar()
	erro := aplicacao.Run(os.Args)

	if erro != nil {
		log.Fatal(erro)
	}
}