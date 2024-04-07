package main

import "fmt"

func main() {
	numero := 10
	if numero > 15 {
		fmt.Println("Maior que 15")
	} else {
		fmt.Println("Menor que 15")
	}

	//If init - é possível declarar uma variável diretamente no IF
	if outroNumero := numero; outroNumero > 0 {
		fmt.Println("Número é maior que zero")
	}
}