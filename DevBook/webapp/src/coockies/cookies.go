package cookies

import (
	"net/http"
	"time"
	"webapp/src/config"

	"github.com/gorilla/securecookie"
)

var s *securecookie.SecureCookie

//Configurar utiliza a variáves de ambiente para a criação do SecureCookie
func Configurar() {
	s = securecookie.New(config.Hashkey, config.BlockKey)
}

func Salvar(w http.ResponseWriter, ID, token string) error{
	dados := map[string]string{
		"id": ID,
		"token": token,
	}

	dadosCodificados, erro := s.Encode("dados", dados)//primeiro parametro é o nome do cookie que se quer dar, o segundo parâmetro são os dados que serão salvos no cookie
	if erro != nil {
		return erro
	}

	http.SetCookie(w, &http.Cookie{
		Name: "dados",//nome do cookie
		Value: dadosCodificados,
		Path: "/",//qual caminho vai ser usado o cookies, no casso usaremos em qualquer lugar na nossa aplicação web por isso passamos a raiz /
		HttpOnly: true,//para evitar que o cookie seja acessado no lado do cliente
	})

	return nil
}

//Ler retorna os valores armazenados no cookie
func Ler(r *http.Request) (map[string]string, error) {
	cookie, erro := r.Cookie("dados")
	if erro != nil {
		return nil, erro
	}
	//alocando um map vazio na memória e guardando na variável
	valores := make(map[string]string)
	if erro = s.Decode("dados", cookie.Value, &valores); erro != nil {
		return nil, erro
	}

	return valores, nil
}

//Deletar remove os valores armazenados no cookie
func Deletar (w http.ResponseWriter) {
	http.SetCookie(w, &http.Cookie{
		Name: "dados",//nome do cookie
		Value: "",//passando os dados do cookie vazio para poder fazer logoff 
		Path: "/",//qual caminho vai ser usado o cookies, no casso usaremos em qualquer lugar na nossa aplicação web por isso passamos a raiz /
		HttpOnly: true,//para evitar que o cookie seja acessado no lado do cliente
		Expires: time.Unix(0,0), //tempo de expiração do cookie, passando como 0 para indicar que o cookie está expirado
	})
}