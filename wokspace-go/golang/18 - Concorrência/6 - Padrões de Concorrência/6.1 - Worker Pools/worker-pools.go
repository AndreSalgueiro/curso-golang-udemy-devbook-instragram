package main

import "fmt"

func main() {
	tarefas := make(chan int, 45)
	resultado := make(chan int, 45)

	//criando vários trabalhadores para pegar as tarefas
	go worker(tarefas,resultado)
	go worker(tarefas,resultado)
	go worker(tarefas,resultado)
	go worker(tarefas,resultado)
	go worker(tarefas,resultado)
	go worker(tarefas,resultado)
	go worker(tarefas,resultado)
	go worker(tarefas,resultado)

	for i := 0; i < 45; i++ {
		tarefas <- i
	}

	for i := 0; i < 45; i++ {
		resultado := <- resultado
		fmt.Println(resultado)
	}
}

func worker(tarefas <-chan int, resultado chan<- int){
	for numero := range tarefas {
		resultado <- fibonacci(numero)
	}
}

func fibonacci(posicao int) int {
	if posicao <= 1 {
		return posicao
	}
	return fibonacci(posicao-2) + fibonacci(posicao-1)
}