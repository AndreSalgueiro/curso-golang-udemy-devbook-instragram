package main

import "fmt"

func main() {
	fmt.Println("Funções Recursivaas")
	//fibonacci
	// 1 1 2 3 5 8 13
	posicao := uint(100)

	for i := uint(0); i < posicao; i++ {
		fmt.Println(fibonacci(i))
	}


}

func fibonacci(posicao uint) uint {
	if posicao <= 1 {
		return posicao
	}
	return fibonacci(posicao-2) + fibonacci(posicao-1)
}