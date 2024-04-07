package main

import "fmt"

func main() {
	fmt.Println(alunoEstaAprovado(7, 8))	
}
//o que estiver com a cláusula defer (adiar) vai ser executado no final de tudo, logo antes do retorno independente da posição de chamada.
func alunoEstaAprovado(n1, n2 float32) bool {
	defer fmt.Println("Média Calculada. Resultado será retornado")
	fmt.Println("Entrando na função que calcula média")

	media := (n1 + n2) / 2

	if media >= 6 {
		return true
	}
	return false
}