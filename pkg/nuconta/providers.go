package nuconta

import (
	"github.com/angelokurtis/money/pkg/nuconta/transactions"
	"github.com/angelokurtis/money/pkg/nuconta/transactions/handlers"
	"github.com/google/wire"
)

var Providers = wire.NewSet(
	Commands,
	transactions.NewService,
	handlers.New,
)
