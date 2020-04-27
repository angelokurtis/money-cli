package transactions

import (
	"time"

	"github.com/angelokurtis/money/internal/log"
	"github.com/angelokurtis/money/pkg/nuconta/transactions/handlers"
	"github.com/pkg/errors"
	"google.golang.org/api/gmail/v1"
)

type Service struct {
	messages *gmail.UsersMessagesService
	handler  handlers.Handler
}

func NewService(messages *gmail.UsersMessagesService, handler handlers.Handler) *Service {
	return &Service{messages: messages, handler: handler}
}

func (s *Service) List(after time.Time) ([]*handlers.Transaction, error) {
	log.Debugf("getting all message with nuconta label from %v", after)
	res, err := s.messages.
		List("me").
		Q("label:bancários-nuconta after:" + after.Format("01/02/2006")).
		Do()
	if err != nil {
		return nil, errors.Wrap(err, "failed to list Nuconta transactions")
	}
	messages := res.Messages
	tx := make([]*handlers.Transaction, 0, len(messages))
	for _, message := range messages {
		t, err := s.transaction(message.Id)
		if err != nil {
			return nil, errors.WithMessage(err, "failed to list Nuconta transactions")
		}
		tx = append(tx, t)
	}
	log.Debugf("found %v messages", len(tx))
	return tx, nil
}

func (s *Service) transaction(id string) (*handlers.Transaction, error) {
	msg, err := s.messageDetails(id)
	if err != nil {
		return nil, errors.WithMessage(err, "failed to list Nuconta transactions")
	}
	return s.handler.Handle(msg)
}

func (s *Service) messageDetails(id string) (*gmail.Message, error) {
	msg, err := s.messages.
		Get("me", id).
		Format("minimal").
		Do()
	if err != nil {
		return nil, errors.Wrap(err, "failed to get email details")
	}
	log.Debugf("got details from message %v of %v", id, time.Unix(0, msg.InternalDate*int64(time.Millisecond)))
	return msg, nil
}
