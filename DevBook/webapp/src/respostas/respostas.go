package respostas

import (
	"encoding/json"
	"log"
	"net/http"
)

type ErroAPI struct {
	Erro string `json:"erro"`
}

// JSON retorna uma resposta em formato JSON para a requisição
func JSON(w http.ResponseWriter, statusCode int, dados interface{}){
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	if statusCode != http.StatusNoContent {
		if erro := json.NewEncoder(w).Encode(dados); erro != nil {
			log.Fatal(erro)
		}
	}
}

func TratarStatusCodeErro(w http.ResponseWriter, r *http.Response) {
	var erroAPI ErroAPI
	json.NewDecoder(r.Body).Decode(&erroAPI)
	JSON(w, r.StatusCode, erroAPI)
}