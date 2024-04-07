package app

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

//gerar vai retornar a aplicação de linha de comando pronta para ser executada
func Gerar() *cli.App {
	app := cli.NewApp()
	app.Name = "Aplicação de linha de comando"
	app.Usage = "Busca IPs e nome de servidores na internet"

	//declaração dos comandos e valores passados para a aplicação
	flags := []cli.Flag {
		cli.StringFlag {
			Name: "host",
			Value: "devbook.com.br", //valor default caso não passe parametros
		},
	}

	//configurando as funcionalidades que minha aplicação vai ter
	app.Commands = []cli.Command {
		{
			Name: "ip", //nome da aplicação que vou chamar na linha de comando
			Usage: "Busca IPs de endereços na interner",//descrição do que ela faz
			Flags: flags,
			Action: buscarIps,
		},
		{
			Name: "servidor",
			Usage: "Busca o nome dos servidores na internet",
			Flags: flags,
			Action: buscarServidores,
		},
	}
	return app
}

func buscarIps(c *cli.Context) {
	host := c.String("host")

	ips, erro := net.LookupIP(host)

	if erro != nil {
		log.Fatal(erro)
	}

	for _,ip := range ips {
		fmt.Println(ip)
	}
}

func buscarServidores(c *cli.Context) {
	host := c.String("host")

	servidores, erro := net.LookupNS(host)

	if erro != nil {
		log.Fatal(erro)
	}

	for _,servidor := range servidores {
		fmt.Println(servidor.Host)
	}
}