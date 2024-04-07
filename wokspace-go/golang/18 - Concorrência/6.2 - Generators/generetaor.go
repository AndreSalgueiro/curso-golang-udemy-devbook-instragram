package main

import (
	"fmt"
	"time"
)

func main() {
	canal := escrever("teste")
	
	for i := 0; i < 10; i++ {
		fmt.Println(<-canal)
	}
}
//Genate basicamente é um padrão que encapsula a chamada de uma goroutine num método, assim chamamos a função normalmente na main
func escrever(texto string) <-chan string {
	canal := make(chan string)

	go func() {
		for {
			canal <- fmt.Sprintf("Valor Recebido $s", texto)
			time.Sleep(time.Millisecond * 500)
		}
	}()
	return canal
}