package rotas

import (
	"net/http"
	"webapp/src/middlewares"

	"github.com/gorilla/mux"
)

//Rota representa todas as rotas da Aplicação
type Rota struct {
	URI    string
	Metodo string
	Funcao func(http.ResponseWriter, *http.Request)
	RequerAutenticacao bool
}

//Configurar coloca todas as rotas dentro de router
func Configurar(router *mux.Router) *mux.Router {
	rotas := rotasLogin
	rotas = append(rotas, rotasUsuarios...)//os tres pontinhos faz um loop interno e adiciona todas as rotas existentes de usuario
	rotas = append(rotas, rotaPaginaPrincipal)
	rotas = append(rotas, rotasPublicacoes...)
	rotas = append(rotas, rotaLogout)

	for _, rota := range rotas {
		if rota.RequerAutenticacao {
			router.HandleFunc(rota.URI, middlewares.Logger(middlewares.Autenticar(rota.Funcao))).Methods(rota.Metodo)
		} else {
			router.HandleFunc(rota.URI, middlewares.Logger(rota.Funcao)).Methods(rota.Metodo)
		}
	}
	
	//Aqui configura um server para apontar para o Go onde estão nossos arquivos css e js sem isso os arquivos não são encontrados
	fileServer := http.FileServer(http.Dir("./assets"))
	router.PathPrefix("/assets").Handler(http.StripPrefix("/assets/", fileServer))
	
	return router
}