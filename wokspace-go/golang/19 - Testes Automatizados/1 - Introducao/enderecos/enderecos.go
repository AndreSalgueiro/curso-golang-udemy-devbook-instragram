package enderecos

import (
	"strings"
)
//TipoEndereco verifica se tem um tipo valido e retorna
func TipoEndereco(endereco string) string{
	tiposEnderecosValidos := []string{"Avenida","Rua","Estrada","Rodovia"}
	
	//faz um split da string por espaço e colocatudo em letra minuscula com ToLower()
	tipoEnderecoRecebido := strings.Split(endereco, " ")[0]
	tipoEnderecoRecebidoMinusculo := strings.ToLower(tipoEnderecoRecebido)

	for _, tipoEnderecoValido := range tiposEnderecosValidos {
		tipoEndecoValidoMinusculo := strings.ToLower(tipoEnderecoValido)
		
		if strings.EqualFold(tipoEndecoValidoMinusculo, tipoEnderecoRecebidoMinusculo) {
			return tipoEnderecoValido
		}
	}
	return "Tipo de endereco inválido"

}