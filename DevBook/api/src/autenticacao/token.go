package autenticacao

import (
	"api/src/config"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

//CriarToken retorna um token assinado com as permissões do usuário
func CriarToken(usuarioID uint64) (string, error) {
	permissoes := jwt.MapClaims{}
	permissoes["authorized"] = true
	permissoes["exp"] = time.Now().Add(time.Hour * 6).Unix()//dizendo que o token expira em 6 horas
	permissoes["usuarioId"] = usuarioID
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, permissoes)//criando o método de assinatura
	return token.SignedString([]byte(config.SecretKey))//assinando o token digitalmente
}

//ValidarToken verifica se o token passado na requisição é válido
func ValidarToken(r *http.Request) error {
	tokenString := extrairToken(r)
	token, erro := jwt.Parse(tokenString, retornaChaveVerificacao)//parse para poder extrair os dados do token
	if erro != nil {
		return erro
	}
	//fmt.Println(token)
	//verifica se existe os Claims (campos do token) e se o token ainda é válido (se não expirou)
	if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return nil
	}
	
	return errors.New("token inváido")
}

func ExtrairUsuarioID(r *http.Request) (uint64, error) {
	tokenString := extrairToken(r)
	token, erro := jwt.Parse(tokenString, retornaChaveVerificacao)
	if erro != nil {
		return 0, erro
	}

	if permissoes, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		usuarioID, erro := strconv.ParseUint(fmt.Sprintf("%.0f", permissoes["usuarioId"]), 10, 64)
		if erro != nil {
			return 0, erro
		}
		return usuarioID, nil
	}
	return 0, errors.New("token invalido")
}

func extrairToken(r *http.Request) string {
	token := r.Header.Get("Authorization")
	// verifica se o token foi passado (o token chega como uma string no formato "BearerToken qqqngnsn4n566nndnfngf")
	if len(strings.Split(token, " ")) == 2 {
		return strings.Split(token, " ")[1]//retorna o token que está na segunda posição após o split
	}

	return ""
}
//retornaChaveVerificacao verifica se o método de assinatura que estou utilizando (SigningMethodHS256) é da familia HMAC
func retornaChaveVerificacao(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("Método de assinatura inesperado! %v", token.Header["alg"])
	}

	return config.SecretKey, nil
}

