package commands

import (
	"github.com/angelokurtis/money/pkg/nuconta"
	"github.com/angelokurtis/money/pkg/spreadsheet"
	"github.com/urfave/cli"
)

func Join(nuconta nuconta.Command, spreadsheet spreadsheet.Command) []cli.Command {
	cmd := make([]cli.Command, 0, 2)
	cmd = append(cmd, cli.Command(nuconta), cli.Command(spreadsheet))
	return cmd
}
