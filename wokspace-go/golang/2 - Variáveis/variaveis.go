package main

import "fmt"

func main() {
	//duas formas de declarar variável
	var variavel1 string = "variável 1"
	variavel := "variável 2"
	fmt.Println(variavel1)
	fmt.Println(variavel)

	//formas de declarar mais de uma variável
	var(
		variavel3 string = "variavel 3"
		variavel4 string = "variavel 4"
	)

	variavel5, variavel6 := "variavel 5", "variavel 6"
	fmt.Println(variavel3)
	fmt.Println(variavel4)
	fmt.Println(variavel5)
	fmt.Println(variavel6)


//também existem várias maneiras de declarar constantes 
	const const1 string = "constante 1"
	fmt.Println(const1)

	const(
		constante3 string = "constante 3"
		constante4 string = "constante 4"
	)

	fmt.Println(constante3)
	fmt.Println(constante4)

	//forma de inverter o valor da variavel em go
	variavel5, variavel6 = variavel6, variavel5
	fmt.Println("Ecemplo de Inverter variável em go")
	fmt.Println(variavel5)
	fmt.Println(variavel6)

}
