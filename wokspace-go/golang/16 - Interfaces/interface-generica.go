package main

import "fmt"

func main() {
	//como o parametro deste método é uma interface genérica, pode-se enviar qualquer coisa
	interfaceGenerica("string")
	interfaceGenerica(5)
	interfaceGenerica(true)

}
//metodo que recebe uma interface generica/vazia que recebe qualquer coisa
func interfaceGenerica(inter interface{}) {
	fmt.Println(inter)
}