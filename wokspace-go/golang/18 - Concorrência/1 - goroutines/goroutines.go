package main 

import (
	"fmt"
	"time"
)

func main() {
	/*com o comando go na frente das funcs ou metodos, o go não vai aguardar o método terminar para executar a próxima tarefa,
	ele manda executar e enquanto isso passa para a próxima linha a ser executada*/
	go escrever("Olá mundo!")//goroutine
	escrever("Programando em Go!")
}

func escrever(texto string) { 
for {
	fmt.Println(texto)
	time.Sleep(time.Second)
	}
}