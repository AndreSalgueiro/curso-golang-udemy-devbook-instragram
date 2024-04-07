package main

import "fmt"

var n int

/*A função init é executada antes da função main,
muito utilizada para inicializar um setup antes do inicio do programa*/
func init() {
	fmt.Println("Executando a função init")
	n=10 
}

func main() {
	fmt.Println("função main sendo executada")
	fmt.Println(n)
}