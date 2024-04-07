package main

import (
	"fmt"
	"teste/auxiliar"

	"github.com/badoux/checkmail"
)

func main() {
	fmt.Println("Escrevendo do arquivo main")
	auxiliar.Escrever()

	erro := checkmail.ValidateFormat("andrecoordena@gmail.com")
	fmt.Println("Validando e-mail: ", erro)
}