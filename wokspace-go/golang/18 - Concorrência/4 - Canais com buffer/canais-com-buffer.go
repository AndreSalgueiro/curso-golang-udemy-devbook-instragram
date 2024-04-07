package main 

import "fmt"

func main() {
	/*criando canal com buffer (2), nesta caso o a execução só será bloqueada, quando atingir a capacidade do buffer*/
	canal := make(chan string, 2)//2 é a quantidade de valores que está esperando receber
	canal <- "olá Mundo"
	canal <- "Programando em Go"

	mensagem := <-canal
	mensagem2 := <-canal
	//mensagem3 := <-canal
	fmt.Println(mensagem)
	fmt.Println(mensagem2)

	//como o buffer do canal está com valor 2, só esperar trafegar 2 valores pelo canal, se eu tentar receber ou enviar um terceiro canal dá erro
	//fmt.println(mensagem3)
}