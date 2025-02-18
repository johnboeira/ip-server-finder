package app

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

func Gerar() *cli.App {
	app := cli.NewApp()
	app.Name = "App de linha de comando"
	app.Usage = "Busca ips e nome de servidor"

	app.Commands = []cli.Command{
		{
			Name:  "ip",
			Usage: "Busca ips de endereços",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "host",
					Value: "amazon.com",
				},
			},
			Action: buscarIps,
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

	for _, ip := range ips {
		fmt.Println(ip)
	}
}
