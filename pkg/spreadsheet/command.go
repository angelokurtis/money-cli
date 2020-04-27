package spreadsheet

import (
	"github.com/angelokurtis/money/internal/log"
	"github.com/urfave/cli"
)

type Command cli.Command

func Commands(svc *Service) Command {
	return Command{
		Name:      "spreadsheet",
		ShortName: "s",
		Subcommands: []cli.Command{
			{
				Name: "classify",
				Action: func(c *cli.Context) {
					err := svc.Classify()
					if err != nil {
						log.Fatal(err)
					}
				},
			},
			{
				Name: "analytics",
				Action: func(c *cli.Context) {
					err := svc.PlotAnalysis()
					if err != nil {
						log.Fatal(err)
					}
				},
			},
		},
	}
}
