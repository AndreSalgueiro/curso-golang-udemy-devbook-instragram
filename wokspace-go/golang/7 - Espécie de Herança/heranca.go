package main

import "fmt"

type pessoa struct {
	nome string
	sobrenome string
	idade uint8
	altura uint8
}

type estudante struct {
	pessoa
	curso string
	faculdade string
}

func main () {
	fmt.Println("Herança")
	p1 := pessoa{"João", "Pedro", 20, 178}
	estudante1 := estudante{p1, "Engenharia", "USP"}
	fmt.Println(estudante1)
	fmt.Println("Nome estudante: ", estudante1.nome)
	
	//outra forma de criar pessoa direto ao criar o estudante
	estudante2 := estudante{pessoa{"Elisa", "Salgueiro", 1, 40}, "bebê", "N/A"}
	fmt.Println(estudante2)
}