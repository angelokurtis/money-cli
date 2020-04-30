package spreadsheet

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/angelokurtis/money/pkg/spreadsheet/classification"
	"github.com/pkg/errors"
)

type Transaction struct {
	Date        time.Time
	Account     string
	Description string
	Value       float64
	Type        string
	Updated     bool
	TypeAddress string
}

func NewTransaction(index int, row []interface{}) (*Transaction, error) {
	var t1 string
	var t2 string
	var updated bool
	if len(row) > 6 {
		t1 = fmt.Sprintf("%s", row[6])
	}
	desc := fmt.Sprintf("%s", row[2])
	if t1 == "" {
		t2 = classification.Classify(desc)
		updated = t2 != ""
	}
	val, err := NewValue(fmt.Sprintf("%s", row[4]))
	if err != nil {
		return nil, err
	}

	date, err := time.Parse("02/01/2006 15:04:05", fmt.Sprintf("%s", row[0]))
	if err != nil {
		return nil, err
	}

	return &Transaction{
		Date:        date,
		Account:     fmt.Sprintf("%s", row[1]),
		Description: desc,
		Value:       val,
		Type: func() string {
			if updated {
				return t2
			}
			return t1
		}(),
		Updated:     updated,
		TypeAddress: fmt.Sprintf("Extratos!J%d", index),
	}, nil
}

func (t *Transaction) Typed() bool {
	return t.Type != ""
}

func (t *Transaction) Profit() bool {
	switch t.Type {
	case "Salário", "Impostos", "Hora Extra", "Outros Descontos", "Outras Receitas":
		return true
	}
	return false
}

func (t *Transaction) Debit() bool {
	switch t.Type {
	case "Bancários", "Serviços", "Casa", "Extras", "Transporte", "Veículos", "Diversos", "":
		return true
	}
	return false
}

func NewValue(val string) (float64, error) {
	val = strings.Replace(val, " ", "", -1)
	val = strings.Replace(val, "R$", "", 1)
	val = strings.Replace(val, ".", "", 1)
	val = strings.Replace(val, ",", ".", 1)
	s, err := strconv.ParseFloat(val, 64)
	if err != nil {
		return 0, errors.Wrap(err, "failed to read transaction value")
	}
	return s, nil
}
