package main

import (
	"fmt"
	"testes-automatizados/enderecos"
)

func main() {
	tipoEndereco := enderecos.TipoEndereco("Avenida almirante Amorim do Vale")

	fmt.Println("o tipo de endereço é: ",tipoEndereco)
}