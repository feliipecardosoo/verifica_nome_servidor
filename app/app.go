package app

import "github.com/urfave/cli"

// Gerar vai retornar a aplicacao de linha de comando pronta para ser executada
func Gerar() *cli.App {
	app := cli.NewApp()
	app.Name = "Aplicacao de Linha de Comando"
	app.Usage = "Busca IPs e Nomes de Servidor"

	return app
}
