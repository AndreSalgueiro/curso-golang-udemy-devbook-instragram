package middlewares

import (
	"log"
	"net/http"
	cookies "webapp/src/coockies"
)

// Logger escreve informações da requisição no terminal
func Logger(proximaFuncao http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("\n %s %s %s", r.Method, r.RequestURI, r.Host)
		proximaFuncao(w, r)
	}
}

// Autenticar verifica a existência de cookies
func Autenticar(proximaFuncao http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if _, erro := cookies.Ler(r); erro != nil{
			//caso não exista dados de autenticação no cookie redireciona para a tela de login
			http.Redirect(w, r, "/login", 302)
			return
		}
		proximaFuncao(w, r)
	}
}
