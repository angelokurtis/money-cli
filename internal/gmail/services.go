package gmail

import (
	"net/http"

	"github.com/pkg/errors"
	"google.golang.org/api/gmail/v1"
)

func NewService(client *http.Client) (*gmail.Service, error) {
	s, err := gmail.New(client)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create Gmail service")
	}
	return s, nil
}

func NewUsersService(service *gmail.Service) *gmail.UsersService {
	return service.Users
}

func NewUsersMessagesService(users *gmail.UsersService) *gmail.UsersMessagesService {
	return users.Messages
}
