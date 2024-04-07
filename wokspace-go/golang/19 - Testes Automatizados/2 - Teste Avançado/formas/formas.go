package formas

import "math"

//interface para a struct trinagulo e circulo
type Forma interface {
	Area() float64
}

//para usar a interface, as structs precisa implementar o método da interface, ou seja o método area()
type Retangulo struct {
	Altura float64
	Largura float64
}
//método da struct retangulo
func (r Retangulo) Area() float64 {
	return r.Altura * r.Largura
}

type Circulo struct {
	Raio float64
}

//método da struct circulo
func (c Circulo) Area() float64 {
	return math.Pi * (c.Raio * c.Raio)
}
