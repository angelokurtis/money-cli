package handlers

import (
	"strings"
	"time"

	"github.com/angelokurtis/money/internal/log"
	"github.com/pkg/errors"
	"google.golang.org/api/gmail/v1"
)

type payment struct {
	next Handler
}

func (h *payment) Handle(message *gmail.Message) (*Transaction, error) {
	body := message.Snippet
	if strings.HasPrefix(body, "Boleto bancário pago com sucesso") {
		date := time.Unix(0, message.InternalDate*int64(time.Millisecond))
		log.Debugf("running payment handler for message %v of %v", message.Id, date)

		stringValue := getStringInBetween(body, "R$", " ")
		value, err := valueFromString(stringValue)
		if err != nil {
			return nil, err
		}

		to := getStringInBetween(body, "Favorecido", "Vencimento")
		if to == "" {
			return nil, errors.New("failed to get payment destination")
		}

		return &Transaction{
			Date:        date.Format("02/01/2006 15:04:05"),
			Description: "Boleto bancário pago para " + to,
			Value:       value * -1,
		}, nil
	} else if strings.HasPrefix(body, "Conta de consumo paga com sucesso") {
		date := time.Unix(0, message.InternalDate*int64(time.Millisecond))
		log.Debugf("running payment handler for message %v of %v", message.Id, date)

		stringValue := getStringInBetween(body, "R$", " ")
		value, err := valueFromString(stringValue)
		if err != nil {
			return nil, err
		}

		to := getStringInBetween(body, "Favorecido", "Código de barras")
		if to == "" {
			return nil, errors.New("failed to get payment destination")
		}

		return &Transaction{
			Date:        date.Format("02/01/2006 15:04:05"),
			Description: "Boleto bancário pago para " + to,
			Value:       value * -1,
		}, nil
	}
	return h.next.Handle(message)
}
