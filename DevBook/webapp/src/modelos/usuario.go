package modelos

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"
	"webapp/src/config"
	"webapp/src/requisicoes"
)

// Usuario representa uma pessoa utilizando a rede social
type Usuario struct {
	ID          uint64       `json:"id"` //passando o comando omitempty emite o campo quando ele está vazio
	Nome        string       `json:"nome"`
	Nick        string       `json:"nick"`
	Email       string       `json:"email"`
	Senha       string       `json:"senha"`
	CriadoEm    time.Time    `json:"CriadoEm"`
	Seguidores  []Usuario    `json:"seguidores"`
	Seguindo    []Usuario    `json:"seguindo"`
	Publicacoes []Publicacao `json:"publicacoes"`
}

// BuscarUsuarioCompleto faz 4 requisições na API para montar o usuário
func BuscarUsuarioCompleto(usuarioID uint64, r *http.Request) (Usuario, error) {
	canalUsuario := make(chan Usuario)
	canalSeguidores := make(chan []Usuario)
	canalSeguindo := make(chan []Usuario)
	canalPublicacoes := make(chan []Publicacao)

	//executando os métodos utilizando go runtines para serem executaas em paralelo
	go BuscarDadosUsuario(canalUsuario, usuarioID, r)
	go BuscarSeguidores(canalSeguidores, usuarioID, r)
	go BuscarSeguindo(canalSeguindo, usuarioID, r)
	go BuscarPublicacoes(canalPublicacoes, usuarioID, r)

	//criando uma variável para cada resposta do canal
	var (
		usuario     Usuario
		seguidores  []Usuario
		seguindo    []Usuario
		publicacoes []Publicacao
	)

	//4 iteração, uma para cada resposta
	for i := 0; i < 4; i++ {
		select {
		case usuarioCarregado := <-canalUsuario: //pegando o dado do canal
			if usuarioCarregado.ID == 0 { //significa que deu erro pois o canal recebeu um usuário vazio
				return Usuario{}, errors.New("Erro ao buscar o usuário")
			}

			usuario = usuarioCarregado
		
		case seguidoresCarregados := <-canalSeguidores: //pegando o dado do canal
			if seguidoresCarregados == nil { //significa que deu erro pois o canal recebeu um slice de usuários vazio
				return Usuario{}, errors.New("Erro ao buscar seguidores")
			}

			seguidores = seguidoresCarregados	

			case seguindoCarregados := <-canalSeguindo: //pegando o dado do canal
			if seguindoCarregados == nil { //significa que deu erro pois o canal recebeu um slice de usuários vazio
				return Usuario{}, errors.New("Erro ao buscar quem o usuário está seguindo")
			}

			seguindo = seguindoCarregados

		case publicacoesCarregadas := <-canalPublicacoes: //pegando o dado do canal
			if publicacoesCarregadas == nil { //significa que deu erro pois o canal recebeu um usuário vazio
				return Usuario{}, errors.New("Erro ao buscar as publicacoes")
			}

			publicacoes = publicacoesCarregadas

		}

	}

	//juntando todas as informações recebidas pelo canal para montar o perfil do usuário
	usuario.Seguidores = seguidores
	usuario.Seguindo = seguindo
	usuario.Publicacoes = publicacoes

	return usuario, nil
}

//o parametro das quatros funções de canal, só recebe informações e não envia, por isso a seta é ao contrário

// BuscarDadosUsuario chama a API para buscar os dados base do usuário
func BuscarDadosUsuario(canal chan<- Usuario, usuarioID uint64, r *http.Request) {
	url := fmt.Sprintf("%s/usuarios/%d", config.APIURL, usuarioID)
	response, erro := requisicoes.FazerRequisicaoComAutenticacao(r, http.MethodGet, url, nil)
	//se der erro devolvo um usuário vazio pelo canal
	if erro != nil {
		canal <- Usuario{}
		return
	}
	defer response.Body.Close()

	var usuario Usuario
	if erro = json.NewDecoder(response.Body).Decode(&usuario); erro != nil {
		canal <- Usuario{}
		return
	}

	canal <- usuario
}

// BuscarSeguidores chama a API para buscar os seguidores do usuário
func BuscarSeguidores(canal chan<- []Usuario, usuarioID uint64, r *http.Request) {
	url := fmt.Sprintf("%s/usuarios/%d/seguidores", config.APIURL, usuarioID)
	response, erro := requisicoes.FazerRequisicaoComAutenticacao(r, http.MethodGet, url, nil)
	//se der erro devolvo slice de usuário vazio pelo canal
	if erro != nil {
		canal <- nil
		return
	}
	defer response.Body.Close()

	var seguidores []Usuario
	if erro = json.NewDecoder(response.Body).Decode(&seguidores); erro != nil {
		canal <- nil
		return
	}

	//pode acontecer o usuário não ter nenhum seguidor, então precisamos tratrar essa situação passando um slice vazio
	if seguidores == nil { 
		canal <- make([]Usuario, 0)
		return
	}

	canal <- seguidores
}

// BuscarSeguindo chama a API para buscar os usuarios seguidos por um usuário
func BuscarSeguindo(canal chan<- []Usuario, usuarioID uint64, r *http.Request) {
	url := fmt.Sprintf("%s/usuarios/%d/seguindo", config.APIURL, usuarioID)
	response, erro := requisicoes.FazerRequisicaoComAutenticacao(r, http.MethodGet, url, nil)
	//se der erro devolvo slice de usuário vazio pelo canal
	if erro != nil {
		canal <- nil
		return
	}
	defer response.Body.Close()

	var seguindo []Usuario
	if erro = json.NewDecoder(response.Body).Decode(&seguindo); erro != nil {
		canal <- nil
		return
	}

	//pode acontecer o usuário não seguir nenhum usuário, então precisamos tratrar essa situação passando um slice vazio
	if seguindo == nil { 
		canal <- make([]Usuario, 0)
		return
	}

	canal <- seguindo
}

func BuscarPublicacoes(canal chan<- []Publicacao, usuarioID uint64, r *http.Request) {
	url := fmt.Sprintf("%s/usuarios/%d/publicacoes", config.APIURL, usuarioID)
	response, erro := requisicoes.FazerRequisicaoComAutenticacao(r, http.MethodGet, url, nil)
	//se der erro devolvo slice de usuário vazio pelo canal
	if erro != nil {
		canal <- nil
		return
	}
	defer response.Body.Close()

	var publicacoes []Publicacao
	if erro = json.NewDecoder(response.Body).Decode(&publicacoes); erro != nil {
		canal <- nil
		return
	}

	//pode acontecer o usuário não ter nenhum seguidor, então precisamos tratrar essa situação passando um slice vazio
	if publicacoes == nil { 
		canal <- make([]Publicacao, 0)
		return
	}

	canal <- publicacoes
}
