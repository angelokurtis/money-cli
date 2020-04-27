//+build wireinject

package app

import (
	"github.com/angelokurtis/money/internal/commands"
	"github.com/angelokurtis/money/internal/gmail"
	"github.com/angelokurtis/money/internal/gsheets"
	"github.com/angelokurtis/money/internal/http"
	"github.com/angelokurtis/money/pkg/nuconta"
	"github.com/angelokurtis/money/pkg/spreadsheet"
	"github.com/google/wire"
	"github.com/urfave/cli"
)

type App struct {
	*cli.App
}

func newApp(commands []cli.Command) *App {
	app := cli.NewApp()
	app.Name = "Simple Nubank Account CLI"
	app.Usage = "An CLI for getting account statement"
	app.Author = "Tiago Angelo"
	app.Version = "1.0.0"
	app.Commands = commands
	return &App{App: app}
}

func New() (*App, error) {
	wire.Build(
		newApp,
		commands.Providers,
		gmail.Providers,
		gsheets.Providers,
		http.Providers,
		nuconta.Providers,
		spreadsheet.Providers,
	)
	return nil, nil
}
