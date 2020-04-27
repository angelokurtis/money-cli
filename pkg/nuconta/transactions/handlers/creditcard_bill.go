package handlers

import (
	"strings"
	"time"

	"github.com/angelokurtis/money/internal/log"
	"google.golang.org/api/gmail/v1"
)

type creditCardBill struct {
	next Handler
}

func (h *creditCardBill) Handle(message *gmail.Message) (*Transaction, error) {
	body := message.Snippet
	if strings.HasPrefix(body, "Fatura paga com sucesso") {
		date := time.Unix(0, message.InternalDate*int64(time.Millisecond))
		log.Debugf("running credit card bill handler for message %v of %v", message.Id, date)

		stringValue := getStringInBetween(body, "R$", " ")
		value, err := valueFromString(stringValue)
		if err != nil {
			return nil, err
		}

		return &Transaction{
			Date:        date.Format("02/01/2006 15:04:05"),
			Description: "Pagamento de fatura realizado com sucesso",
			Value:       value * -1,
		}, nil
	}
	return h.next.Handle(message)
}
