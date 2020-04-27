package handlers

import (
	"strings"
	"time"

	"github.com/angelokurtis/money/internal/log"
	"google.golang.org/api/gmail/v1"
)

type deposit struct {
	next Handler
}

func (h *deposit) Handle(message *gmail.Message) (*Transaction, error) {
	body := message.Snippet
	if strings.HasPrefix(body, "Você recebeu uma transferência") {
		date := time.Unix(0, message.InternalDate*int64(time.Millisecond))
		log.Debugf("running deposit handler for message %v of %v", message.Id, date)

		stringValue := getStringInBetween(body, "R$", " ")
		value, err := valueFromString(stringValue)
		if err != nil {
			return nil, err
		}

		from := getStringInBetween(body, "transferência de", "e o valor")
		if from == "" {
			from = "origem não identificada"
		}

		return &Transaction{
			Date:        date.Format("02/01/2006 15:04:05"),
			Description: "Transferência recebida de " + from,
			Value:       value,
		}, nil
	} else if strings.HasPrefix(body, "Recebemos sua primeira transferência") || strings.HasPrefix(body, "Recebemos sua transferência") {
		date := time.Unix(0, message.InternalDate*int64(time.Millisecond))
		log.Debugf("running deposit handler for message %v of %v", message.Id, date)

		stringValue := getStringInBetween(body, "R$", " ")
		value, err := valueFromString(stringValue)
		if err != nil {
			return nil, err
		}

		return &Transaction{
			Date:        date.Format("02/01/2006 15:04:05"),
			Description: "Transferência recebida de Tiago de Sousa Angelo",
			Value:       value,
		}, nil
	}
	return h.next.Handle(message)
}
