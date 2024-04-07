package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
)

type cachorro struct {
	Nome  string `json: "nome"`
	Raca  string `json: "raca"`
	Idade uint `json: "idade"`
}

func main() {
	c := cachorro{"Rex", "Dalmatas", 3}
	fmt.Print("\n Cachorro como struct ", c)

	//transformando uma struct em json
	cJSON, erro := json.Marshal(c)
	if erro != nil {
		log.Fatal(erro)
	}
	//vai imprimir o cachorro formatado como jsom mais com valores em bytes para formatar como json precisa utilizar outra função de transformação
	fmt.Println("\n Cachorro como Json de bytes: ", cJSON)
	//Para formatar a apresentação de um json
	fmt.Println("\n Cachorro como Json: ", bytes.NewBuffer(cJSON))

	//dá para transformar um map em json também
	c2 := map[string]string {
		"nome": "Toby",
		"raca": "Poodle",
	}

	c2JSON, erro := json.Marshal(c2)
	if erro != nil {
		log.Fatal(erro)
	}

	//vai imprimir o cachorro formatado como jsom mais com valores em bytes para formatar como json precisa utilizar outra função de transformação
	fmt.Println("\n Cachorro como Json de bytes: ", c2JSON)
	//Para formatar a apresentação de um json
	fmt.Println("\n Cachorro como Json: ", bytes.NewBuffer(c2JSON))

	
}