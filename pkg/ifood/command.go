package ifood

import (
	"fmt"
	"github.com/angelokurtis/money/pkg/ifood/deliveries"
	"github.com/olekukonko/tablewriter"
	"os"
	"time"

	"github.com/angelokurtis/money/internal/log"
	"github.com/pkg/errors"
	"github.com/urfave/cli"
)

type Command cli.Command

func Commands(deliveries *deliveries.Service) Command {
	return Command{
		Name: "ifood",
		Subcommands: []cli.Command{
			{
				Name: "deliveries",
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
					reqs, err := deliveries.List(after)
					if err != nil {
						log.Fatal(err)
					}
					table := tablewriter.NewWriter(os.Stdout)
					table.SetBorder(false)
					table.SetAutoWrapText(false)
					table.SetHeader([]string{"Data", "Local", "Pedido", "Valor"})
					for _, req := range reqs {
						table.Append([]string{req.Datetime.String(), req.Place, req.Id, fmt.Sprintf("%.2f", req.Value)})
					}
					table.Render()
				},
				Flags: []cli.Flag{
					cli.StringFlag{
						Name:  "after",
						Usage: "initial deliveries date",
					},
				},
			},
		},
	}
}
