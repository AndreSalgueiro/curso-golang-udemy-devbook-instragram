package main

import (
	"log"
	"net/http"
	"text/template"
)
//variável utilizada para renderilar uma página HTML
var templates *template.Template

type usuario struct {
	Nome string
	Email string
}

func main() {
	//criando um padrão de arquivo para ele reconhecer, nesta caso .html
	templates = template.Must(template.ParseGlob("*.html"))

	//criando uma rota de chamada no servidor de requisição. A função func() sempre vai receber esses dois parâmetros
	http.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {

		//criando um usuário para passar para a página
		u := usuario{"João","joaopedro@gmail.com"}

		/*para executar a pagina html basta passar os parâmetros w - response, nome da página que será renderizada
		e último parâmetro seria algum dado que gostaríamos passar para a página html*/
		templates.ExecuteTemplate(w, "home.html", u)		
	})

	//subindo um servidor no go. A chamada precisa estár na última linha pois depois dessa chamda o programa trava
	log.Fatal(http.ListenAndServe(":5000", nil))
}
