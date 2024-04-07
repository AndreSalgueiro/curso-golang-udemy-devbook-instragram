package main

import "fmt"

func main() {
	soma, subtracao := calcuclosMatemáticos(10,20)
	fmt.Println(soma,subtracao)
}
//retorno nomeado
func calcuclosMatemáticos(n1,n2 int) (soma int, subtracao int) {
//não precisa declarar a variavel pois já foi mencionada no retorno
soma = n1 + n2
subtracao = n1 - n2
//não precisa passar o nome das variável de retorno
return
}