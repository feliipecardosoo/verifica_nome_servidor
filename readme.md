# Retorna IP e Nome Servidor
Este projeto permite obter informações de IP e o nome do servidor hospedado de um site específico, fornecido como parâmetro.

## Como executar
Para executar este projeto localmente, é necessário ter o Go instalado.


1. Clone este repositório:

2. Copile o projeto:

```bash
go build -o modulo.exe
```


3. Execute o seguinte comando para iniciar o projeto:

**Para retornar o IP do site:**
```bash
./modulo.exe ip --host nomesite.com.br
```

**Para retornar o nome do servidor hospedado:** 
```bash
./modulo.exe servidores --host nomesite.com.br
```

## Estratégia 
As funções principais deste projeto estão nos arquivos: app/buscarIps e app/buscarServidores.

**Exemplo do arquivo app/buscarServidores:**
```bash
func buscarServidores(c *cli.Context) {
	host := c.String("host")

	servidores, erro := net.LookupNS(host)
	if erro != nil {
		log.Fatal(erro)
	}

	for _, servidor := range servidores {
		fmt.Println(servidor.Host)
	}
}

```
A função buscarServidores é responsável por receber uma requisição, verificar um possível erro e retornar uma lista com os nomes dos servidores.

## Dependências
- **[Go (versão 1.16 ou superior)](https://go.dev/dl/)**
- **[urfave/cli (biblioteca para construção de interfaces de linha de comando)](https://github.com/urfave/cli)**
