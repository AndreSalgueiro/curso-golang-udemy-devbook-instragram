package main

import (
	"log"
	"net/http"
)

func main() {
	//criando uma rota de chamada no servidor de requisição. A função func() sempre vai receber esses dois parâmetros
	http.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Olá mundo!"))
	})

	

	//outra forma de passar a função no método é declarar fora e passar somente o nome dela
	http.HandleFunc("/usuario", usuario)

	//subindo um servidor no go. A chamada precisa estár na última linha pois depois dessa chamda o programa trava
	log.Fatal(http.ListenAndServe(":5000", nil))
}

func usuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Acessando dados do usuario"))
}