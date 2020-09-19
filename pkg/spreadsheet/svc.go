package spreadsheet

import (
	"fmt"
	"sort"
	"time"

	"github.com/angelokurtis/money/internal/log"
	"github.com/pkg/errors"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
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
	b, err := c.balances()
	if err != nil {
		return err
	}
	balances := make([]*Balance, 0, 0)
	for _, current := range b {
		if current.Date.After(time.Date(2018, 12, 01, 0, 0, 0, 0, time.UTC)) {
			balances = append(balances, current)
		}
	}

	profit := make(plotter.Values, 0, len(balances))
	debit := make(plotter.Values, 0, len(balances))
	names := make([]string, 0, len(balances))
	for _, balance := range balances {
		p := balance.Profit / 1000
		if p < 0 {
			p = 0
		}
		profit = append(profit, p)
		d := balance.Debit / 1000
		if d > 0 {
			d = 0
		}
		debit = append(debit, d*-1)
		names = append(names, balance.Date.Format("Jan/2006"))
	}

	p, err := plot.New()
	if err != nil {
		return err
	}
	p.Title.Text = "Balanço mensal"
	p.Y.Label.Text = "R$ (x1000)"

	w := vg.Points(20)

	barsA, err := plotter.NewBarChart(profit, w)
	if err != nil {
		return err
	}
	barsA.LineStyle.Width = vg.Length(0)
	barsA.Color = plotutil.Color(1)
	barsA.Offset = -w

	barsB, err := plotter.NewBarChart(debit, w)
	if err != nil {
		return err
	}
	barsB.LineStyle.Width = vg.Length(0)
	barsB.Color = plotutil.Color(0)

	p.Add(barsA, barsB)
	p.Legend.Add("Créditos", barsA)
	p.Legend.Add("Débitos", barsB)
	p.Legend.Top = true
	p.NominalX(names...)
	if err := p.Save(3*5*vg.Inch, 3*3*vg.Inch, "balance.png"); err != nil {
		return err
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
