package main

import (
	"fmt"
	"time"
)

func main() {
	//multiplexador junta a saída de dois ou mais canais em apenas 1 canal
	canalSaida := multiplexador(escrever("Olá mundo"), escrever("Programando em Go"))

	for i := 0; i<10; i++ {
		fmt.Println(<-canalSaida)
	}
	
}

func multiplexador(canalEntrada1, CanalEntrada2 <-chan string) <-chan string {
	canalSaida := make(chan string)

	go func() {
		for {
			select {
			case mensagem := <-canalEntrada1:
				canalSaida <- mensagem
			case mensagem := <-CanalEntrada2:
				canalSaida <- mensagem
			}
		}
	}()
	return canalSaida
} 

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