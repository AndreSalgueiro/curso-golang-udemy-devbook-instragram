package router

import (
	"api/src/router/rotas"

	"github.com/gorilla/mux"
)

//Gerar retorna uma nova rota
func Gerar() *mux.Router {
	r := mux.NewRouter()
	return rotas.Configurar(r)
}