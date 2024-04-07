package main

import "fmt"

func main() {
	numero := 20
	/*quando passamos um parâmetro para a função, é feito uma cópia
	  da variável e a alteração que é feita na variável dentro da função
	  não afeta a variável que está fora como nesse exemplo*/	
	numeroInvertidoValor := inverteSinalValor(numero)
	fmt.Println("Variavel retornado da funcao: ",numeroInvertidoValor)
	fmt.Println("Variavel de fora da funcao: ",numero)
	
	fmt.Println("-------------------------------------------------------------")
	
	/*quando passamos um ponteiro como parâmetro para a função, estamos passando o endereço de memória 
	onde está variável está, logo a alteração no valor armazenado neste endereço de memória feita 
	no dentro da função também afeta o valor da variável que está fora desta função como no exemplo
	*/
	novoNumero := 40
	fmt.Println("Variável de fora da função: ",novoNumero)
	inverteSinalReferencia(&novoNumero)
	fmt.Print("Variável de fora da função depois de ser alterada pela função: ",novoNumero)
}
func inverteSinalValor(numero int) int {
	return numero * -1
}
func inverteSinalReferencia(novoNumero *int){
	*novoNumero = *novoNumero * -1
}