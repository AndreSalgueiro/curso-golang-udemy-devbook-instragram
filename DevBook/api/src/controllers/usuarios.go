package controllers

import (
	"api/banco"
	"api/src/repositorios"
	seguranca "api/src/Seguranca"
	"api/src/autenticacao"
	"api/src/modelos"
	"api/src/respostas"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

func CriarUsuario(w http.ResponseWriter, r *http.Request) {
	corpoRequest, erro := io.ReadAll(r.Body)
	if erro != nil {
		//log.Fatal(erro)
		respostas.Erro(w, http.StatusUnprocessableEntity, erro)
		return
	}
	
	var usuario modelos.Usuarios
	if erro =json.Unmarshal(corpoRequest, &usuario); erro != nil {
		//log.Fatal(erro)
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	if erro = usuario.Preparar("cadastro"); erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}
	
	db, erro := banco.Conectar()
	if erro != nil {
		//log.Fatal(erro)
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repositorio := repositorios.NovoRepositorioUsuario(db)
	usuario.ID, erro = repositorio.Criar(usuario)
	if erro != nil {
		//log.Fatal(erro)
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	//retorna os dados do usuário inserido caso tudo deu certo
	respostas.JSON(w, http.StatusCreated, usuario)
}
//BuscarUsuarios busca um usuário pelo nome ou nick
func BuscarUsuarios(w http.ResponseWriter, r *http.Request) {
	nomeNick := strings.ToLower(r.URL.Query().Get("usuario"))//na query passada na url pega o valor do paramentro usuarios
	db, erro := banco.Conectar() 
	if erro != nil{
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repositorio := repositorios.NovoRepositorioUsuario(db)
	usuarios, erro := repositorio.Buscar(nomeNick)
	if erro != nil {
		respostas.Erro(w,http.StatusInternalServerError,erro)
		return
	}
	respostas.JSON(w,http.StatusOK, usuarios)
}
func BuscarUsuario(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)

	usuarioID, erro := strconv.ParseInt(parametros["usuarioId"], 10, 64)
	if erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repositorio := repositorios.NovoRepositorioUsuario(db)
	usuario, erro := repositorio.BuscarPorID(usuarioID)
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
	}

	respostas.JSON(w, http.StatusOK, usuario)
}
func AtualizarUsuario(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)
	usuarioID, erro := strconv.ParseUint(parametros["usuarioId"], 10, 64)
	if erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	usuarioIDToken, erro := autenticacao.ExtrairUsuarioID(r)
	if erro != nil {
		respostas.Erro(w, http.StatusUnauthorized, erro)
		return
	}

	if usuarioID != usuarioIDToken {
		respostas.Erro(w, http.StatusForbidden, errors.New("não é possível atualizar um usuário que não seja o seu"))
		return
	}

	corpoRequest, erro := io.ReadAll(r.Body)
	if erro != nil {
		//log.Fatal(erro)
		respostas.Erro(w, http.StatusUnprocessableEntity, erro)
		return
	}
	
	var usuario modelos.Usuarios
	if erro =json.Unmarshal(corpoRequest, &usuario); erro != nil {
		//log.Fatal(erro)
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	if erro = usuario.Preparar("edicao"); erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}
	
	db, erro := banco.Conectar()
	if erro != nil {
		//log.Fatal(erro)
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repositorio := repositorios.NovoRepositorioUsuario(db)
	erro = repositorio.Atualizar(usuarioID, usuario)
	if erro != nil {
		//log.Fatal(erro)
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	//retorna os dados do usuário inserido caso tudo deu certo
	respostas.JSON(w, http.StatusNoContent, nil)
}

func DeletarUsuario(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)
	usuarioID, erro := strconv.ParseUint(parametros["usuarioId"], 10, 64)
	if erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	usuarioIDToken, erro := autenticacao.ExtrairUsuarioID(r)
	if erro != nil {
		respostas.Erro(w, http.StatusUnauthorized, erro)
		return
	}

	if usuarioID != usuarioIDToken {
		respostas.Erro(w, http.StatusForbidden, errors.New("não é possível deletar um usuário que não seja o seu"))
		return
	}
	
	db, erro := banco.Conectar()
	if erro != nil {
		//log.Fatal(erro)
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repositorio := repositorios.NovoRepositorioUsuario(db)
	erro = repositorio.Deletar(usuarioID)
	if erro != nil {
		//log.Fatal(erro)
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	//retorna os dados do usuário inserido caso tudo deu certo
	respostas.JSON(w, http.StatusNoContent, nil)
}
//SeguirUsuario permite que um usuario siga outro
func SeguirUsuario(w http.ResponseWriter, r *http.Request) {
	
	seguidorID, erro := autenticacao.ExtrairUsuarioID(r)
	if erro != nil {
		respostas.Erro(w, http.StatusUnauthorized, erro)
		return
	}

	parametros := mux.Vars(r)
	usuarioID, erro := strconv.ParseUint(parametros["usuarioId"], 10, 64)
	if erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	if seguidorID == usuarioID {
		respostas.Erro(w, http.StatusForbidden, errors.New("não é possível seguir você mesmo"))
		return
	}
	
	db, erro := banco.Conectar()
	if erro != nil {
		//log.Fatal(erro)
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repositorio := repositorios.NovoRepositorioUsuario(db)
	if erro = repositorio.Seguir(usuarioID, seguidorID); erro != nil{
		//log.Fatal(erro)
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	respostas.JSON(w, http.StatusNoContent, nil)
}

//PararSeguirUsuario permite que um usuario pare de seguir o outro
func PararSeguirUsuario(w http.ResponseWriter, r *http.Request) {
	
	seguidorID, erro := autenticacao.ExtrairUsuarioID(r)
	if erro != nil {
		respostas.Erro(w, http.StatusUnauthorized, erro)
		return
	}

	parametros := mux.Vars(r)
	usuarioID, erro := strconv.ParseUint(parametros["usuarioId"], 10, 64)
	if erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	if seguidorID == usuarioID {
		respostas.Erro(w, http.StatusForbidden, errors.New("não é possível deixar de seguir você mesmo"))
		return
	}
	
	db, erro := banco.Conectar()
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repositorio := repositorios.NovoRepositorioUsuario(db)
	if erro = repositorio.PararSeguir(usuarioID, seguidorID); erro != nil{
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	respostas.JSON(w, http.StatusNoContent, nil)
}

//BuscarSeguidores busca todos os seguidores de um usuário
func BuscarSeguidores(w http.ResponseWriter, r *http.Request) {
	
	parametros := mux.Vars(r)
	usuarioID, erro := strconv.ParseUint(parametros["usuarioId"], 10, 64)
	if erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}
	
	db, erro := banco.Conectar()
	if erro != nil {
		//log.Fatal(erro)
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repositorio := repositorios.NovoRepositorioUsuario(db)
	seguidores, erro := repositorio.BuscarSeguidores(usuarioID)
	if erro != nil{
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	respostas.JSON(w, http.StatusOK, seguidores)
}

//BuscarSeguindo busca todos os usuários que um determinado usuário está seguindo
func BuscarSeguindo(w http.ResponseWriter, r *http.Request) {
	
	parametros := mux.Vars(r)
	usuarioID, erro := strconv.ParseUint(parametros["usuarioId"], 10, 64)
	if erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}
	
	db, erro := banco.Conectar()
	if erro != nil {
		//log.Fatal(erro)
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repositorio := repositorios.NovoRepositorioUsuario(db)
	seguidores, erro := repositorio.BuscarSeguindo(usuarioID)
	if erro != nil{
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	respostas.JSON(w, http.StatusOK, seguidores)
}

//AtualizarSenha permite atualizar a senha do usuário
func AtualizarSenha(w http.ResponseWriter, r *http.Request) {
	usuarioIDToken, erro := autenticacao.ExtrairUsuarioID(r)
	if erro != nil {
		respostas.Erro(w, http.StatusUnauthorized, erro)
		return
	}
	
	parametros := mux.Vars(r)
	usuarioID, erro := strconv.ParseUint(parametros["usuarioId"], 10, 64)
	if erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	if usuarioID != usuarioIDToken {
		respostas.Erro(w, http.StatusForbidden, errors.New("não é possível atualizar a senha de um usuário que não seja o seu"))
		return
	}

	corpoRequest, erro := io.ReadAll(r.Body)
	if erro != nil {
		//log.Fatal(erro)
		respostas.Erro(w, http.StatusUnprocessableEntity, erro)
		return
	}
	
	var senha modelos.Senha
	if erro = json.Unmarshal(corpoRequest, &senha); erro != nil {
		//log.Fatal(erro)
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		//log.Fatal(erro)
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repositorio := repositorios.NovoRepositorioUsuario(db)
	senhaSalvaBanco, erro := repositorio.BuscarSenha(usuarioID)
	if erro != nil {
		//log.Fatal(erro)
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	if erro = seguranca.VerificaSenha(senhaSalvaBanco, senha.Atual); erro != nil {
		respostas.Erro(w, http.StatusUnauthorized, errors.New("a senha atual não condiz com a senha cadastrada"))
	}

	senhaComHash, erro := seguranca.Hash(senha.Nova)
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	if erro = repositorio.AtualizarSenha(usuarioID, string(senhaComHash)); erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	respostas.JSON(w, http.StatusNoContent, nil)
}