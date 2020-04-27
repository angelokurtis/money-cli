package handlers

import (
	"google.golang.org/api/gmail/v1"
)

type Handler interface {
	Handle(message *gmail.Message) (*Transaction, error)
}

func New() Handler {
	return &creditCardBill{
		next: &deposit{
			next: &payment{
				next: &transfer{
					next: &defaultH{},
				},
			},
		},
	}
}
