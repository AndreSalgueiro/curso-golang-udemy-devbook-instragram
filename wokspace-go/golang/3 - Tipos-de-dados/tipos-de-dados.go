package main

import (
	"errors"
	"fmt"
)

func main() {
	//tipos de numeros inteiros
	var number int8 = 100
	var number2 int16 = 100
	var number3 int32 = 100
	var mumber4 int64 = 100
	fmt.Println(number,"\n",number2,"\n",number3,"\n",mumber4)
	
	//declara um numero inteiro po inferencia, nesse caso terá o tamanho máximo de acordo com a arquitetura da máquina onde o programa está rodando.
	var number5 int = 100
	number6 := 100
	fmt.Println(number5,"\n",number6)

	//tipo de inteiro sem sinal, ou seja, não aceita números negativos
	var unumber uint8 = 100
	var unumber2 uint16 = 100
	var unumber3 uint32 = 100
	var umumber4 uint64 = 100 
	fmt.Println(unumber,"\n",unumber2,"\n",unumber3,"\n",umumber4)

	//alias para int32 = rune
	var number7 rune = 100
	fmt.Println(number7)

	//Alias para unt8 = byte
	var number8 byte = 100
	fmt.Println(number8)

	//numeros flutuantes
	var fnumber float32 = 1.00
	var fnumber2 float64 = 100.00
	fmt.Println(fnumber,"\n",fnumber2)

	//float com inferencia. O tamanho vai ser de acordo com a arquitetura da máquina
	fnumber3 := 100.00
	fmt.Println(fnumber3)

	//char é declarado como aspas simples e o que retorna é o número representante na tabela ASK
	//Não exite char, o que retorna é um int representando o numero da tabela ASK
	char := 'B'
	fmt.Println(char)

	//bolenao
	var booleano bool = false
	fmt.Println(booleano)

	//error no go é um tipo de dados
	var erro error 
	var erro2 error = errors.New("Setando uma mensagem de erro")
	fmt.Println(erro,"\n", erro2)
	
}