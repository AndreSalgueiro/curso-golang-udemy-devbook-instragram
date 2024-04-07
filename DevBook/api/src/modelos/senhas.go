package modelos

//senha representa o formato da requisicao de senha
type Senha struct {
	Nova string `json:"nova"`
	Atual string `json:"atual"`
}