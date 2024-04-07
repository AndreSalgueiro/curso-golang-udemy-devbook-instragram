package enderecos_test

import (
	. "testes-automatizados/enderecos"//esse ponto espaço na frente do import é para criar um alias, assim não precisa colocar o nome do pacote na frente da chamada da função
	"testing"
)

type cenarioTeste struct {
	enderecoInserido string
	retornoEsperado string
}

//TESTE DE UNIDADE
// a função tem que começar com a palavra em imglês "test" e o nome do arquivo.go tem que ter o padrão de nome "_test.go"
//toda função de teste recebe o parâmetro do tipo ponteiro "*testing.T"
func TestTipoDeEndereco(t *testing.T){
	t.Parallel() //indica que esse teste pode rodar em paralelo com outros, precisa incluir em cada teste que vai rodar em paralelo
	//criando um slice de struct
	cenarioTeste := []cenarioTeste {
		//endereçoInserido, retorno esperado
		{ "Rua abc", "Rua" },
		{ "Rodovia abc", "Rodovia" },
		{ "Avenida abc", "Avenida" },
		{ "Estrada abc", "Estrada" },
		{ "Mercaso abc", "Tipo de endereco inválido" },
		{ "", "Tipo de endereco inválido" },
		{ "AVENIDA abc", "Avenida" },
	}
	
	for _, cenario := range cenarioTeste {
		retornoRecebido := TipoEndereco(cenario.enderecoInserido)

		if retornoRecebido != cenario.retornoEsperado {
			t.Errorf("O tipo recebido %s é diferente do esperado %s,", retornoRecebido, cenario.retornoEsperado)
		}
	}

}
