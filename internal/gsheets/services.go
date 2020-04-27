package gsheets

import (
	"net/http"

	"github.com/pkg/errors"
	"google.golang.org/api/sheets/v4"
)

func NewService(client *http.Client) (*sheets.Service, error) {
	s, err := sheets.New(client)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create Google Sheets service")
	}
	return s, nil
}
