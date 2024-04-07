package main 

import "fmt"

func main() {
	total := soma(2,3,5,1,7,9,0,1,3,200)
	fmt.Println(total)
	fmt.Println("-----------------------------------")

	escrever("Testando parametros de função", 6,2,6,44,22,33,4,0,8,7,)
}
//declarando números "infinitos" de parametros na função
func soma(numeros ...int) (total int){

	for _,numero := range numeros {
		total += numero
	}
	return 
}

//é possível receber um parâmetro e mais infinitos parametros. Obs o parametro variatico precisa ser o último parametro
//Não é possível passar duas variáveis com parametros infinitos na função
func escrever(frase string, numeros ...int) {
	for _, numero := range numeros {
		fmt.Println(frase, numero)
	}
	
}