package main

import (
	"strings"
	"os"
	"github.com/urfave/cli/v2"
	"github/http"
	"github/internal/handlers"
) 

func main()  {
	
	app := &cli.App {
		Name: "github-client",
		Usage: "Fetch data of github with the API",
		Action: func (ctx *cli.Context) error {
			var name = ctx.Args().First()

			if (strings.TrimSpace(name) == "") {
				panic("Invalid username!")
			}

			events, err := http.FetchEvents(name)

			if (err != nil) {
				panic(err)
			}

			format := handlers.Format {
				Events: &events, 
			}

			format.PrintEvents()

			return nil
		},
	}

	if err := app.Run(os.Args); err != nil { 
		panic(err)	
	}

}

