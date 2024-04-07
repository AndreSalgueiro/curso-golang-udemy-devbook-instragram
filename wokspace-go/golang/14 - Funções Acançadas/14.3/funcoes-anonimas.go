package main

import "fmt"

func main() {
	//funcao anonima. o "()" no final da função é como deve ser invocada uma função anônima
	teste := "teste"
	func(texto string) {
		fmt.Println(texto)
		fmt.Println(teste)
	}("Olá Mundo")

	//funcao anonima com retorno de parametro
	retorno := func(texto string) string {
		return fmt.Sprintf("Parametro recebido -> %s %d", texto, 10)
	}("Olá mundo")

	fmt.Println(retorno)
}
