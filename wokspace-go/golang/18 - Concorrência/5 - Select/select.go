package main

import (
	"time"
	"fmt"
)

func main() {
	//criando 2 canais
	canal1, canal2 := make(chan string), make(chan string)

	//criando duas goroutines

	go func() {
		for {
			time.Sleep(time.Millisecond * 500)//meio segundos
			canal1 <- "Canal 1"
		}
	}()

	go func() {
		for {
			time.Sleep(time.Second * 2)
			canal2 <- "Canal 2"
		}
	}()

	/*se usar um for infinito para esperar os valores chegarem pelo canal, não é o ideal
	pois o programa fica bloqueado esperando um dado chegar em um canal para depois esperar o dado chegar no segundo canal*/
	for {
		/*vai ficar parado aqui até receber um dado para depois esperar o dado do próximo canal, mesmo as funções 
		que enviam informações pelo canal serem uma goroutine*/
		/*mensagem1 := <- canal1
		fmt.Println(mensagem1)
	
		mensagem2 := <- canal2
		fmt.Println(mensagem2)*/

		//para resolver a situação acima utilizamos o select assim não fica bloqueado esperando um dado chegar no canal
		select {
		case mensagem1 := <- canal1:
			fmt.Println(mensagem1)
		case mensagem2 := <- canal2:
		fmt.Println(mensagem2)
		}
	}


}
