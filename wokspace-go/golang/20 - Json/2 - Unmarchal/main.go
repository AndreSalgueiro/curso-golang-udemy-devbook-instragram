package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type cachorro struct {
	Nome  string `json: "nome"`
	Raca  string `json: "raca"`
	Idade uint   `json: "idade"`
}

func main() {
	//criando uma string no formato Json
	cJSON := `{"Nome":"Rex","Raca":"Dalmatas","Idade":3}`

	//criando um cachorro vazio
	var c cachorro
	/*convertendo um Json cJSON numa struct c. Na conversão precisa passar um array de bytes para a função Unmarshal() 
	e também o endereço de memória da variável cachorro para o conteúdo dela poder ser alterado, no caso de vazio para os valores do Json*/
	if erro := json.Unmarshal([]byte(cJSON), &c); erro != nil {
		log.Fatal(erro)
	}
	fmt.Println(c)

	//Agora transformando um Json num Map
	c2JSON := `{"nome":"Toby","raca":"Poodle"}`

	c2 := make(map[string]string)

	if erro := json.Unmarshal([]byte(c2JSON), &c2); erro != nil {
		log.Fatal(erro)
	}
	fmt.Println(c2)

}