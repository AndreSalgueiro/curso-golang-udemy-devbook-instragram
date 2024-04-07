package formas

import (
	"math"
	"testing"
)

func TestArea(t *testing.T) {
	//o método Run() é utilizado para deixar mais organizado a separação de teste de duas funcções
	t.Run("Retângulo", func(t *testing.T) {
		ret := Retangulo{10,12}
		areaEsperada := float64(120)
		areaRecebida := ret.Area()
		if areaEsperada != areaRecebida {
			t.Errorf("Área recebida %f é diferente da esperada %f", areaRecebida, areaEsperada)
		}
	})

	t.Run("Circulo", func(t *testing.T) {
		circ := Circulo {10}
		areaEsperado := float64(math.Pi * 100)
		areaRecebida := circ.Area()
		if areaEsperado != areaRecebida {
			t.Errorf("Área recebida %f é diferente da esperada %f", areaRecebida, areaEsperado)
		}

	})
}