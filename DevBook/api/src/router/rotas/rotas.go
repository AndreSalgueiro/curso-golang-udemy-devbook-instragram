package rotas

import (
	"api/src/middlewares"
	"net/http"

	"github.com/gorilla/mux"
)

//Rota representa todas as rotas da API
type Rota struct {
	URI    string
	Metodo string
	Funcao func(http.ResponseWriter, *http.Request)
	RequerAutenticacao bool
}
//Configurar coloca todas as rotas dentro do Router
func Configurar(r *mux.Router) *mux.Router {
	rotas := rotasUsuarios //adicionando as rotas de usuário
	rotas = append(rotas, rotaLogin)//adicionando a rota de login no slice
	rotas = append(rotas, rotasPublicacoes...)//os três pontinhos é para ele entender que tem que fazer um append para cada rota configurada no slice de rotas de publicações

	for _, rota := range rotas {
		if rota.RequerAutenticacao {
			r.HandleFunc(rota.URI, middlewares.Logger(middlewares.Autenticar(rota.Funcao))).Methods(rota.Metodo)
		} else {
			r.HandleFunc(rota.URI, middlewares.Logger(rota.Funcao)).Methods(rota.Metodo)
		}
	}
	return r
}