package app

import (
	"github.com/urfave/cli"
)

// Gerar vai retornar a aplicacao de linha de comando pronta para ser executada
func Gerar() *cli.App {
	app := cli.NewApp()
	app.Name = "Aplicacao de Linha de Comando"
	app.Usage = "Busca IPs e Nomes de Servidor"

	flags := []cli.Flag{
		cli.StringFlag{
			Name:  "host",
			Value: "github.com",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:   "ip",
			Usage:  "Busca IPS de enderecos na internet",
			Flags:  flags,
			Action: buscarIps,
		},
		{
			Name:   "servidores",
			Usage:  "Busca o nome dos servidores na internet",
			Flags:  flags,
			Action: buscarServidores,
		},
	}

	return app
}
