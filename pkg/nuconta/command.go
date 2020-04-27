package nuconta

import (
	"encoding/json"
	"io/ioutil"
	"time"

	"github.com/angelokurtis/money/internal/log"
	"github.com/angelokurtis/money/pkg/nuconta/transactions"
	"github.com/pkg/errors"
	"github.com/urfave/cli"
)

type Command cli.Command

func Commands(transactions *transactions.Service) Command {
	return Command{
		Name:      "nuconta",
		ShortName: "nu",
		Subcommands: []cli.Command{
			{
				Name: "statement",
				Action: func(c *cli.Context) {
					var after time.Time
					arg := c.String("after")
					if arg != "" {
						after2, err := time.Parse("02/01/2006", arg)
						after = after2
						if err != nil {
							log.Fatal(errors.Wrap(err, "failed to read param as date"))
						}
					} else {
						after = time.Now()
					}
					messages, err := transactions.List(after)
					if err != nil {
						log.Fatal(err)
					}
					file, _ := json.MarshalIndent(messages, "", " ")
					_ = ioutil.WriteFile("extrato-nuconta.json", file, 0644)
				},
				Flags: []cli.Flag{
					cli.StringFlag{
						Name:  "after",
						Usage: "initial transactions date",
					},
				},
			},
		},
	}
}
