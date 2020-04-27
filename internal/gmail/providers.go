package gmail

import (
	"github.com/google/wire"
)

var Providers = wire.NewSet(
	NewService,
	NewUsersService,
	NewUsersMessagesService,
)
