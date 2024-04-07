package main

import (
	"crud/servidor"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	/*mux é utilizado para gerenciar as funções 
	READ -> GET
	POST -> CREATE 
	DELETE -> DELETE
	UPDATE -> PUT
	*/
	router := mux.NewRouter()
	//criando rota
	router.HandleFunc("/usuarios", servidor.CriarUsuario).Methods(http.MethodPost)
	router.HandleFunc("/usuarios", servidor.BuscarUsuarios).Methods(http.MethodGet)
	router.HandleFunc("/usuarios/{id}", servidor.BuscarUsuario).Methods(http.MethodGet)
	router.HandleFunc("/usuarios/{id}", servidor.AtualizarUsuario).Methods(http.MethodPut)
	router.HandleFunc("/usuarios/{id}", servidor.DeletarUsuario).Methods(http.MethodDelete)

	fmt.Println("Executando na porta 5000")
	//subindo um servidor na porta 5000
	log.Fatal(http.ListenAndServe(":5000", router))
}