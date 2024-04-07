package repositorios

import (
	"api/src/modelos"
	"database/sql"
	"fmt"
)

type Usuarios struct {
	db *sql.DB
}

func NovoRepositorioUsuario(db *sql.DB) *Usuarios {
	return &Usuarios{db}
}
//Criar isere um usuário no banco de dados 
func (repositorio Usuarios) Criar(usuario modelos.Usuarios) (uint64, error) {
	statement, erro := repositorio.db.Prepare("insert into usuarios (nome, nick, email, senha) values (?,?,?,?)",)
	if erro != nil {
		return 0, erro
	}
	defer statement.Close()

	resultado, erro := statement.Exec(usuario.Nome, usuario.Nick, usuario.Email, usuario.Senha)
	if erro != nil {
		return 0, erro
	}
	
	ultimoIdInserido, erro := resultado.LastInsertId()
	if erro != nil {
		return 0, erro
	}

	return uint64(ultimoIdInserido), nil
}
//Buscar traz todos os usuários cadastrados na base
func (repositorio Usuarios) Buscar(nomeNick string) ([]modelos.Usuarios, error) {
	nomeNick = fmt.Sprintf("%s%%",nomeNick)//nomeNick% 
	linhas, erro := repositorio.db.Query("select id,nome,nick,email,criadoEm from usuarios where nome LIKE ? or nick LIKE ?", nomeNick,nomeNick)
	
	if erro != nil {
		return nil, erro
	}
	defer linhas.Close()

	var usuarios []modelos.Usuarios

	for linhas.Next() {
		var usuario modelos.Usuarios
		if erro = linhas.Scan(&usuario.ID, &usuario.Nome, &usuario.Nick, &usuario.Email, &usuario.CriadoEm); erro != nil {
			return nil, erro
		}

		usuarios = append(usuarios, usuario)
	}
	return usuarios, nil
}	
//BuscarPorID busca o usuário pelo id na base de dados
func (repositorio Usuarios) BuscarPorID(ID int64) (modelos.Usuarios, error) {
	linha, erro := repositorio.db.Query("select id, nome, nick, email, criadoEm from usuarios where id = ?", ID)
	if erro != nil {
		//se deu erro devolvo um usuário em branco pois não aceita retorna nil
		return modelos.Usuarios{}, erro
	}
	defer linha.Close()

	var usuario modelos.Usuarios

	if linha.Next() {
		if erro = linha.Scan(
			&usuario.ID,
			&usuario.Nome,
			&usuario.Nick,
			&usuario.Email,
			&usuario.CriadoEm,
		); erro != nil {
			return modelos.Usuarios{}, erro
		}
	}
	return usuario, nil
}
//Atualizar altera um usuário na base de dados
func (repositorio Usuarios) Atualizar(ID uint64, usuario modelos.Usuarios) (error){
	statement, erro := repositorio.db.Prepare(
		"update usuarios set nome =?, nick = ?, email = ? where id = ?",
	)
	if erro != nil {
		return erro
	}
	defer statement.Close()

	_, erro = statement.Exec(usuario.Nome, usuario.Nick, usuario.Email, ID)
	if erro != nil {
		return erro
	}

	return nil
}
//Deletar exclui as informações no banco e dados
func (repositorio Usuarios) Deletar(ID uint64) (error) {
	statement, erro := repositorio.db.Prepare(
		"delete from usuarios where id = ?",
	)
	if erro != nil {
		return erro
	}
	defer statement.Close()

	_, erro = statement.Exec(ID)
	if erro != nil {
		return erro
	}

	return nil
}

func (repositorio Usuarios) BuscarPorEmail(email string) (modelos.Usuarios, error) {
	linha, erro := repositorio.db.Query("select id, senha from usuarios where email =?", email)
	if erro != nil {
		return modelos.Usuarios{}, erro
	}
	defer linha.Close()

	var usuario modelos.Usuarios
	if linha.Next() {
		if erro = linha.Scan(&usuario.ID, &usuario.Senha); erro != nil {
			return modelos.Usuarios{}, erro
		}
	}
	return usuario, nil
}

//Seguir permite que um usuário siga outro
func (repositorio Usuarios) Seguir(usuarioID, seguidorID uint64) (error) {
	//a clousula ignore no insert, vai ignorar caso já exista essa primary+key composta no banco, isso evita de dar erro de violação de constraint
	statement, erro := repositorio.db.Prepare("insert ignore into seguidores (usuario_id, seguidor_id) values (?,?)")
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro = statement.Exec(usuarioID, seguidorID); erro != nil {
		return erro
	}

	return nil
}
//PararSeguir permite que um usuário pare de seguir o outro
func (repositorio Usuarios) PararSeguir(usuarioID, seguidorID uint64) (error){
	statement, erro := repositorio.db.Prepare(
		"delete from seguidores where usuario_id = ? and seguidor_id = ?",
	)
	if erro != nil {
		return erro
	}
	defer statement.Close()

	_, erro = statement.Exec(usuarioID, seguidorID)
	if erro != nil {
		return erro
	}

	return nil
}

//BuscarSeguidores traz todos os seguidores de um usuario
func (repositorio Usuarios) BuscarSeguidores(usuarioID uint64) ([]modelos.Usuarios, error){
	
	linhas, erro := repositorio.db.Query(`
		select u.id, u.nome, u.nick, u.email, u.criadoEm
		from usuarios u inner join seguidores s on u.id = s.seguidor_id where s.usuario_id = ?
	`, usuarioID)
	
	if erro != nil {
		return nil, erro
	}
	defer linhas.Close()

	var usuarios []modelos.Usuarios

	for linhas.Next() {
		var usuario modelos.Usuarios
		if erro = linhas.Scan(&usuario.ID, &usuario.Nome, &usuario.Nick, &usuario.Email, &usuario.CriadoEm); erro != nil {
			return nil, erro
		}

		usuarios = append(usuarios, usuario)
	}
	return usuarios, nil
}

//BuscarSeguindo traz todos os usuários que um determinado usuário está seguindo
func (repositorio Usuarios) BuscarSeguindo(usuarioID uint64) ([]modelos.Usuarios, error){
	
	linhas, erro := repositorio.db.Query(`
		select u.id, u.nome, u.nick, u.email, u.criadoEm
		from usuarios u inner join seguidores s on u.id = s.usuario_id where s.seguidor_id = ?
	`, usuarioID)
	
	if erro != nil {
		return nil, erro
	}
	defer linhas.Close()

	var usuarios []modelos.Usuarios

	for linhas.Next() {
		var usuario modelos.Usuarios
		if erro = linhas.Scan(&usuario.ID, &usuario.Nome, &usuario.Nick, &usuario.Email, &usuario.CriadoEm); erro != nil {
			return nil, erro
		}

		usuarios = append(usuarios, usuario)
	}
	return usuarios, nil
}
//BuscarSenha tras a senha do usuário pelo ID
func (repositorio Usuarios) BuscarSenha(usuarioID uint64) (string, error) {
	linha, erro := repositorio.db.Query("select senha from usuarios where id =?", usuarioID)
	if erro != nil {
		return "", erro
	}
	defer linha.Close()

	var usuario modelos.Usuarios
	if linha.Next() {
		if erro = linha.Scan(&usuario.Senha); erro != nil {
			return "", erro
		}
	}
	return usuario.Senha, nil
}

//AtualizarSenha altera a senha de um usuário na base de dados
func (repositorio Usuarios) AtualizarSenha(usuarioID uint64, senha string) (error){
	statement, erro := repositorio.db.Prepare(
		"update usuarios set senha =? where id = ?",
	)
	if erro != nil {
		return erro
	}
	defer statement.Close()

	_, erro = statement.Exec(senha, usuarioID)
	if erro != nil {
		return erro
	}

	return nil
}