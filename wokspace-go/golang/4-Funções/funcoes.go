package main

import "fmt"

//retorna um numero inteiro
func somar(n1 int8, n2 int8) int8 {
	return n1+n2
}

//dois retornos na funcao
func calculosMatematicos(n1, n2 int8) (int8, int8) {
	soma := n1+n2
	subtracao := n1-n2
	return soma, subtracao
}

//Exemplo de funçao retornado outra função
func multiplica(n1, n2 int) (int, func()) {
	resultado := n1*n2
	imprimir := func() {
		fmt.Println(resultado)
	}
	return resultado, imprimir
}

func main () {
	soma := somar(10, 10)
	fmt.Println(soma)

	//declarar uma variavel do tipo função
	var imprimir = func(paramentro int8) {
		fmt.Println("Imprimindo parametro...: ",paramentro)
	}
	imprimir(soma)

	//obtendo 2 retornos da funcao
	soma, subtracao := calculosMatematicos(10, 15)
	fmt.Println(soma, subtracao)

	//obtendo somente 1 retorno numa funcao que retorna 2 parametros
	somar, _ := calculosMatematicos(10, 15)
	fmt.Println(somar)

	//obtendo o resultado de uma função que retorna outra função
	_, impressao := multiplica(2, 4)
	impressao()
}
