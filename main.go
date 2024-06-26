package main

import (
	"log"
	"modulo/app"
	"os"
)

func main() {

	aplicacao := app.Gerar()
	if erro := aplicacao.Run(os.Args); erro != nil {
		log.Fatal(erro)
	}
}
