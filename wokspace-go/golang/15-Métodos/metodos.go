package main

import "fmt"

type usuario struct {
	nome string
	idade uint8
}
//criando um metodo para uma struct usuario. Agora poddo dizer que a struct tem um método chamado salvar()
func (u usuario) salvar() {
	fmt.Printf("Salvando os dados do Usuário %s no banco de dados\n", u.nome)
}
// dá para retornar valores em métodos
func (u usuario) maiorDeIdade() bool {
	return u.idade >= 18
}

//utilizando ponteiros em metodos
func (u *usuario) fazerAniversario() {
	u.idade++
}

func main() {
	usuario1 := usuario{"Elisa ", 0}
	fmt.Println(usuario1)
	usuario1.salvar()

	usuario2 := usuario{"Davi", 3}
	usuario2.salvar()

	maiorIdade := usuario2.maiorDeIdade()
	fmt.Println(maiorIdade)

	usuario1.fazerAniversario()
	fmt.Println(usuario1.idade)
}