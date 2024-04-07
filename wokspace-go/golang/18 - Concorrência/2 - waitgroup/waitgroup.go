package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	//waitgroup monitora as goroutines e só encerra o programa quando todas elas terminarem
	var waitGroup sync.WaitGroup

	//adiciona a quantidade de goroutines que serão monitoradas e que terá que esperar terminar
	waitGroup.Add(2)

	//cria-se uma go routine com função anônima e coloca a chamada da nossa função dentro dela, uma para cada função
	go func ()  {
		escrever("olá Mundo!")
		waitGroup.Done()//decrementa a quantidade de goroutime que falta executar
	}()

	go func ()  {
		escrever("programando em Go")
		waitGroup.Done()
	}()
	/*Essa chamada faz com que o programa fique nesse ponto aguandando as goroutines terminarem, 
	ou seja,só vai executar as próximas linhas quando terminar a execução das goroutines*/

	waitGroup.Wait()
	
	//só vai executar após todas as goroutines ter terminado
	fmt.Println("não sou uma goroutine")
	
}

func escrever(texto string) { 
for i := 0; i<5;i++{
	fmt.Println(texto)
	time.Sleep(time.Second)
	}
}