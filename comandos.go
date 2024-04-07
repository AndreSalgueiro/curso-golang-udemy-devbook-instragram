go build
go build -v # buildar de forma verbosa
go run <nome do arquivo> #rodar o programa
go mod init modulo # cria o arquivo que vai conter a declaração de todos os packages utilizados no projeto
go get github.com/badoux/checkmail #baixa o pacote e adiciona a dependência no gerenciador de dependência .mod (executar o comando no nível onde está o .mod para funcionar)
go mod tidy #remove todas as dependências não utilizadas no arquivo .mod
gofmt -w ./internal #paara formatar o código
go test # executa os casos de teste que foram escritos no package que vc está executando
go test -v # executa os casos de teste que foram escritos no package que vc está executando de forma verbosa
go test ./... executa o teste de todos os pacotes. (./... significa entrar em todos os pacotes e subpacotes)
go test --cover #vai indicar a porcentagem de cobertura de testes do código
go test --coverprofile cobertura.txt #cria um arquivo chamado cobertura que indica quais códigos não foram cobertos pelo teste
go tool cover --func=cobertura.txt #interpreta o relatório de cobertura de testes e printa na tela, é mais fácil para entender o que não está coberto pelos testes
go tool cover --html=cobertura.txt #mostra uma pagina html indicando exatamente o código que foi coberto e o código que não foi coberto pelos testes.
#Adicionar um gerenciador de espaço de trabalho
cd $WORK #entrar no diretório raiz
go work init #criar o go.work na raiz do projeto
go work use ./tools/ ./tools/gopls/ #adicionar os subdiretórios que contem o go.mod
Digitar html:5 o editor vai autocompletar com uma estrutura básica de html

Pacotes exencias
go get github.com/joho/godotenv #Para manipular variáveis de ambiente
github.com/go-sql-driver/mysql #drive de conexão com banco de dados mysql

############################################ATALHOS#####################################################################
ctrl + shift + l #ao selecionar a palavra, vai alterar todas as linhas que contem a palavra de uma só vez