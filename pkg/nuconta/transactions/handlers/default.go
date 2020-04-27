package handlers

import (
	"fmt"
	"time"

	"github.com/pkg/errors"
	"google.golang.org/api/gmail/v1"
)

type defaultH struct {
	next Handler
}

func (h *defaultH) Handle(message *gmail.Message) (*Transaction, error) {
	return nil, errors.New(fmt.Sprintf("It was not found a handler for message %v of %v", message.Id, time.Unix(0, message.InternalDate*int64(time.Millisecond))))
}
