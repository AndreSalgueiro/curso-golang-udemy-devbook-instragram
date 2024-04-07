package main 

import "fmt"

func main() {
	fmt.Println(alunoEstaAprovado(6,2))
	fmt.Println("Pós execução")
}

func alunoEstaAprovado(n1, n2 float32) bool {
	//essa função é executada no final de tudo inclusive após execução do panic
	//mesmo se não chegar a executar o panic, o defer é executado mais o recover() vai ser nulo então não faz nada.
	defer recuperarExecuxao()
	
	media := (n1 + n2) /2

	if media > 6 {
		return true
	} else if media < 6 {
		return false
	}
	//cai aqui se a média é igual a 6
	//se o programa chegar a executar o panic, o programa para de executar imediatamente chama todas as funções que tem default
	//é diferente do error, pois qnado o programa dá erro, vc faz  tratamento e o programa continua sua execução
	panic("A mádia e exatamente 6!")
}

func recuperarExecuxao() {
	//a função recover() é usada para verificar se houve panic, ou seja se o retorno dela é diferente de null
	if r := recover(); r != nil {
		fmt.Println("Execução recuperada com sucesso")
	}
}