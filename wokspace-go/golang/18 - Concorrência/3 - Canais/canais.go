package main

import (
	"fmt"
	"time"
)

func main() {
	//criando um canal para enviar e receber dados de uma goroutines
	canal := make(chan string) //criando um canal que trafega string

	go escrever("olá Mundo", canal)

	fmt.Println("Depois da função escrever começar a ser executada")

	/*fica em loop esperando uma informação chegar pelo canal, nesse momento
	o programa fica parado aguardando a informação chegar*/
	for {
		mensagem, aberto := <- canal //recebe o valor do canal e verifica se ainda está aberto
		if !aberto {
			break
		}
		fmt.Println(mensagem)
	}

	/*outra forma de fazer um loop esperando os dados chegar no canal, neste caso não precisa verificar
	se o cnaal está fechado, vai encerrar quando terminar de receber os dados*/
	for mensagem  := range canal {
		fmt.Println(mensagem)
	}

	fmt.Println("Fim do Programa")
}

func escrever(texto string, canal chan string) {
	for i := 0; i < 5; i++ {
		canal <- texto //pega o valor e envia para o canal
		time.Sleep(time.Second)
	}
	//após terminar de enviar os valores, precisa fechar o canal
	close(canal)
}