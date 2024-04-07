package main

import "fmt"

func main() {

	//declarando um map, primeira chave é a chave e fora é o valor
	// as chaves precisam ser do mesmo tipo. Os valores precisam ser do mesmo tipo
	usuario := map[string]string {
		"nome": "Pedro",
		"sobrenome": "Silva",
	}
	fmt.Println(usuario["nome"])

	//map aninhado
	usuario2 := map[string]map[string]string{
		"nome": {
			"primeiro": "João",
			"ultimo": "Pedro",
		},
		"curso": {
			"nome": "Engenharia",
			"campus": "Campus 1",
		},
	}
	fmt.Println("----------------------------------------------------------------")
	fmt.Println(usuario2["nome"])
	fmt.Println(usuario2["nome"]["primeiro"])

	//Excluir um valor do map basta passar o nome do map e a chave do valor a ser excluido
	delete(usuario2,"nome")

	//adicionar uma chave ao map depois dele criado
	usuario2["signo"] = map[string]string {
		"nome": "gemeos",
	}
	fmt.Println("--------------------------------------------------------------")
	fmt.Println(usuario2)


}