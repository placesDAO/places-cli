package main

import (
	"os"

	"github.com/urfave/cli"
)

func main() {
	tokenId := ""
	showSVG := false
	showJson := false
	app := &cli.App{
		Name:  "PlacesDAO CLI",
		Usage: "Explore the PlacesDAO smart contract",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "token",
				Value: "1",
				Usage: "the token id to display",
			},
			&cli.BoolFlag{
				Name:  "show",
				Usage: "display the SVG of the token",
			},
			&cli.BoolFlag{
				Name:  "json",
				Usage: "display the json of the token",
			},
		},
		Action: func(c *cli.Context) error {
			tokenId = c.String("token")
			showSVG = c.Bool("show")
			showJson = c.Bool("json")
			return nil
		},
	}

	err := app.Run(os.Args)
	Ok(err)

	token := getToken(tokenId)

	token.print()

	if showJson {
		token.createJson()
	}

	if showSVG {
		token.displayToken()
	}

}
