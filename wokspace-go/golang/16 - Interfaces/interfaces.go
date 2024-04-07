package main

import (
	"fmt"
	"math"
)

func main() {
	r := retangulo{10,5}
	escreverArea(r)

	c := circulo{4}
	escreverArea(c)
	
}

//interface para a struct trinagulo e circulo
type forma interface {
	area() float64
}

//para usar a interface, as structs precisa implementar o método da interface, ou seja o método area()
type retangulo struct {
	altura float64
	largura float64
}
//método da struct retangulo
func (r retangulo) area() float64 {
	return r.altura + r.largura
}

type circulo struct {
	raio float64
}

//método da struct circulo
func (c circulo) area() float64 {
	return math.Pi * (c.raio * c.raio)
}

func escreverArea(f forma) {
	fmt.Println("a área é: ", f.area())
}
