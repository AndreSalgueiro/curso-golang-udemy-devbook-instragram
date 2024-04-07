package main

import "fmt"

func main() {
	//############################### ARRAY #######################################
	//declara um array de strings de 5 posições
	var array1 [5]string
	//incluir uma string na posicao 1
	array1[0] = "Posicao 1"
	fmt.Println(array1)

	//outra forma de criar array
	array2 := [5]string{"posicao 1"}
	fmt.Println(array2)

	//cria um arrai com o tamanho de acordo com a quantidade de dados
	array3 := [...]int{1,2,3,4,5}
	fmt.Println(array3)
	//############################## SLICE ###########################################
	//slice é um tipo de array que tem o tamanho dinâmico, é só não passar o tamanho dele que já é declaro um slice
	slice := []int{1,2,3}
	fmt.Println(slice)
	//adiciona mais um item ao slice, retorna um slice novo contendo o item
	slice = append(slice, 18)
	fmt.Println(slice)

	//dá para pegar um pedaço de um array e salvar dentro de um slice
	slice2 := array3[1:3]
	fmt.Println(slice2)
	//array na verdade é um ponteiro para os valores na memória
	array3[1] = 9
	fmt.Println(slice2)

	fmt.Println("-----------------------------------------------------------")

	//Arrays internos
	//reservando espaço de me mória de 10 posições e no máquino 15 posições para um slice
	//se for atingindo a capacidade automaticamente o Go vai dobrar a capacidade do slice e assim por diante
	slice3 := make([]float32, 10, 15)
	//também dá para criar um slice sem informar a capacidade
	slice4 := make([]float32, 10)
	fmt.Println(slice3)
	fmt.Println(len(slice3))
	fmt.Println(cap(slice3))
	fmt.Println("------------------------------------------------------------")
	fmt.Println(slice4)
	fmt.Println(len(slice4))
	fmt.Println(cap(slice4))
}