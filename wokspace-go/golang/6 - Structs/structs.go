package main

import "fmt"

type usuario struct {
	nome string
	idade uint8
	endereco endereco
}

type endereco struct {
	logradoro string
	numero int8
}

func main () {
	fmt.Println("Arquivo structs")
	var u usuario 
	u.nome = "André"
	u.idade = 21

	fmt.Println(u)

	enderecoEx := endereco{"Rua dos Bobos", 0}

	//outro jeito de criar variável do tipo struct
	usuario2 := usuario{"Elisa", 10, enderecoEx}
	fmt.Println(usuario2)

	//informado um valor do campo da struct
	usuario3 := usuario{idade: 22}
	fmt.Println(usuario3)
}