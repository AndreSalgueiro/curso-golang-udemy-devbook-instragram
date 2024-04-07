package main 

import "fmt"

func main(){
	texto := "imprimindo variável dentro da main"
	fmt.Println(texto)

	funcaoNova := closure()
	funcaoNova()
}
//função que retorna outra função
func closure() func() {
	texto := "imprimendo a variável dentro da funcao closure"

	funcao := func() {
		fmt.Println(texto)
	}
	return funcao
}
//outro jeito de retornar uma função com retorno nomeado
func closure2() (funcao func()) {
	texto := "dentro da funcao closure"

	funcao = func() {
		fmt.Println(texto)
	}
	return
}