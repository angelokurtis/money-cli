package ifood

import (
	"github.com/angelokurtis/money/pkg/ifood/deliveries"
	"github.com/google/wire"
)

var Providers = wire.NewSet(
	Commands,
	deliveries.NewService,
)
