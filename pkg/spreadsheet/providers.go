package spreadsheet

import (
	"github.com/google/wire"
)

var Providers = wire.NewSet(
	Commands,
	NewService,
)
