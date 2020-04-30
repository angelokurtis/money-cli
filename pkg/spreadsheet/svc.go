package spreadsheet

import (
	"fmt"
	"sort"
	"time"

	"github.com/angelokurtis/money/internal/log"
	"github.com/pkg/errors"
	"google.golang.org/api/sheets/v4"
)

const spreadsheetId = "158wWmEHpuWNFf2xaM2bZLZiLuEG0hIxTDNJMXMnXhAk"

type Service struct {
	svc *sheets.Service
}

func NewService(svc *sheets.Service) *Service {
	return &Service{svc: svc}
}

func (c *Service) Classify() error {
	txs, err := c.transactions()
	if err != nil {
		return err
	}
	for _, tx := range txs {
		if tx.Updated {
			val := tx.Type
			var vr sheets.ValueRange
			vr.Values = append(vr.Values, []interface{}{val})
			_, err = c.svc.Spreadsheets.Values.Update(spreadsheetId, tx.TypeAddress, &vr).ValueInputOption("RAW").Do()
			if err != nil {
				return errors.Wrap(err, "unable to update data to sheet")
			}
			log.Debugf("%s was updated with %s", tx.TypeAddress, val)
		}
	}
	return nil
}

func (c *Service) PlotAnalysis() error {
	balances, err := c.balances()
	if err != nil {
		return err
	}
	for _, current := range balances {
		log.Debugf("%s", current.String())
	}
	return nil
}

func (c *Service) transactions() ([]*Transaction, error) {
	initialRow := 2
	sheetRange := fmt.Sprintf("Extratos!D%d:J", initialRow)
	res, err := c.svc.Spreadsheets.Values.Get(spreadsheetId, sheetRange).Do()
	if err != nil {
		return nil, errors.Wrap(err, "unable to retrieve data from sheet")
	}

	txs := make([]*Transaction, 0, len(res.Values))
	for n, row := range res.Values {
		tx, err := NewTransaction(n+initialRow, row)
		if err != nil {
			return nil, err
		}
		txs = append(txs, tx)
	}
	return txs, nil
}

func (c *Service) balances() ([]*Balance, error) {
	txs, err := c.transactions()
	if err != nil {
		return nil, err
	}
	months := make(map[string]*Balance, 0)
	for _, tx := range txs {
		month := tx.Date.Format("01/2006")
		b := months[month]
		if b == nil {
			date, err := time.Parse("01/2006", month)
			if err != nil {
				return nil, err
			}
			b = &Balance{Date: date}
		}
		if tx.Profit() {
			b.Profit += tx.Value
		}
		if tx.Debit() {
			b.Debit += tx.Value
		}
		months[month] = b
	}

	balances := make([]*Balance, 0, len(months))
	for _, v := range months {
		balances = append(balances, v)
	}
	sort.Slice(balances, func(i, j int) bool {
		return balances[i].Date.Before(balances[j].Date)
	})
	for i, current := range balances {
		if i > 0 {
			last := balances[i-1]
			current.Accumulated += current.Profit + current.Debit + last.Accumulated
		} else {
			current.Accumulated += current.Profit + current.Debit
		}
		log.Debugf("%s", current.String())
	}
	return balances, nil
}
