package main

import (
	"fmt"
	"time"
)

func main() {
	i := 0

	for i < 10 {
		i++
		fmt.Println("Incrementando i: ", i)
		time.Sleep(time.Second)
	}

	fmt.Println("--------------------------------------------------------------------")

	//incrementado de 2 em 2
	for j := 0; j< 10; j += 2 {
		fmt.Println("Incrementando j:" , j)
		time.Sleep(time.Second)
	}

	fmt.Println("--------------------------------------------------------------------")

	//RANGE utilizado quando se quer fazer iteração pelos índices
	nomes := [3]string{"João", "Davi", "Lucas"}
	//range nomes - interage pelo array nomes. i - recebe o índice. valor - recebe o valor do índice
	for i, valor := range nomes {
		fmt.Println(i, valor)
	}

	fmt.Println("--------------------------------------------------------------------")

	//se quiser só o indice pode fazer assim
	nomes2 := [3]string{"João", "Davi", "Lucas"}
	for i := range nomes2 {
		fmt.Println(i)
	}

	fmt.Println("--------------------------------------------------------------------")

	//se quiser pegar somente o valor pode fazer assim
	nomes3 := [3]string{"João", "Davi", "Lucas"}
	for _, valor := range nomes3 {
		fmt.Println(valor)
	}

	//Interar por uma string. não colocar a variavel "valor" dentro da string, vali imprimir o código ASK da letra ao invés da letra
	for i, valor := range "PALAVRA" {
		fmt.Println(i, string(valor))
	}

	//Interar pelo MAP
	usuario := map[string]string {
		"nome":	"Leonardo",
		"sobrenome":	"Dias",
	}
	for chave, valor := range usuario {
		fmt.Println(chave, valor)
	}

	//Loop infinito
	for {
		fmt.Println("Executando infinitamente")
		time.Sleep(time.Second)
	}
}
