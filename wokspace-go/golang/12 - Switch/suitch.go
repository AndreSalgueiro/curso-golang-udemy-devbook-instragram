package main 

import "fmt"

func main() {

}

func diaDaSemana(numero int) string {
	switch numero {
	case 1:
		return "Domingo"
	case 2 :
		return "Segunda-feira"
	case 3 : 
		return "Terça-feira"
	default: 
		return "Numero invalido"
	}
}	
//Outra forma de declarar switch
func diaDaSemana2(numero int) string {
	switch {
	case numero == 1:
		return "Segunda-feira"
	case numero == 2:
		return "Terça-feira"
	default:
		return "Numero inválido"
	}
}
//Outra forma de fazer
func diaDaSemana3(numero int) string {
	var diaDaSemana string
	
	switch {
	case numero == 1:
		diaDaSemana = "Segunda-feira"
	case numero == 2:
		diaDaSemana = "Terça-feira"
	default:
		diaDaSemana = "Numero inválido"
	}
	return diaDaSemana
}
//parametro fallthrough
func diaDaSemana4(numero int) string {
	var diaDaSemana string
	
	switch {
	case numero == 1:
		diaDaSemana = "Segunda-feira"
		//atribui o resultado desta condição com o valor da próxima confição
		fallthrough
	case numero == 2:
		diaDaSemana = "Terça-feira"
	default:
		diaDaSemana = "Numero inválido"
	}
	return diaDaSemana
}

