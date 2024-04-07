package main

import "fmt"

func main () {
	//declaração comum de variável
	var var1 int
	//declaração de um ponteiro de inteiro
	var ponteiro *int
	// ponteiro tem o valor default nil, inteiro tem 0
	fmt.Println(ponteiro)
	fmt.Println(var1)

	//para atribuir o valor de uma variável a um ponteiro, precisa por o "&" na frente da variável para ela atribuir o endereço da variável e não o valor
	ponteiro = &var1

	var1 = 10 
	//para imprimir o conteúdo de um ponteiro precisa colocar "*" senão bai imprimir o endereço de memória
	fmt.Println(*ponteiro)
	fmt.Println(var1)
}