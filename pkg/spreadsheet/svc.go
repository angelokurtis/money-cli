package spreadsheet

import (
	"fmt"

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
	initialRow := 2
	sheetRange := fmt.Sprintf("Extratos!D%d:J", initialRow)
	res, err := c.svc.Spreadsheets.Values.Get(spreadsheetId, sheetRange).Do()
	if err != nil {
		return errors.Wrap(err, "unable to retrieve data from sheet")
	}

	if len(res.Values) == 0 {
		log.Debug("No data found.")
	} else {
		for n, row := range res.Values {
			tx, err := NewTransaction(row)
			if err != nil {
				return err
			}
			if tx.Updated {
				val := tx.Type
				var vr sheets.ValueRange
				vr.Values = append(vr.Values, []interface{}{val})
				_, err = c.svc.Spreadsheets.Values.Update(spreadsheetId, fmt.Sprintf("Extratos!J%d", n+initialRow), &vr).ValueInputOption("RAW").Do()
				if err != nil {
					return errors.Wrap(err, "unable to update data to sheet")
				}
				log.Debugf("Extratos!J%d was updated with %s", n+initialRow, val)
			}
		}
	}
	return nil
}

func (c *Service) PlotAnalysis() error {
	initialRow := 2
	sheetRange := fmt.Sprintf("Extratos!D%d:J", initialRow)
	res, err := c.svc.Spreadsheets.Values.Get(spreadsheetId, sheetRange).Do()
	if err != nil {
		return errors.Wrap(err, "unable to retrieve data from sheet")
	}

	txs := make([]*Transaction, 0, len(res.Values))
	if len(res.Values) == 0 {
		log.Debug("No data found.")
	} else {
		for n, row := range res.Values {
			tx, err := NewTransaction(row)
			if err != nil {
				return err
			}
			txs = append(txs, tx)
			log.Debugf("Extratos!J%d = %s R$ %.2f", n+initialRow, tx.Description, tx.Value)
		}
	}
	return nil
}
