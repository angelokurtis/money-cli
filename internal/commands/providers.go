package commands

import (
	"github.com/google/wire"
)

var Providers = wire.NewSet(Join)
