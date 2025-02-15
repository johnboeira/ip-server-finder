package app

import "github.com/urfave/cli"

func Gerar() *cli.App {
	app := cli.NewApp()
	app.Name = "App de linha de comando"
	app.Usage = "Busca ips e nome de servidor"

	return app
}
