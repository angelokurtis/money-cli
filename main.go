package main

import (
	"os"

	"github.com/angelokurtis/money/internal/log"
	"github.com/angelokurtis/money/pkg/app"
)

func main() {
	a, err := app.New()
	if err != nil {
		log.Fatal(err)
	}
	err = a.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
