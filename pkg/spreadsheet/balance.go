package spreadsheet

import (
	"fmt"
	"time"
)

type Balance struct {
	Date        time.Time
	Profit      float64
	Debit       float64
	Accumulated float64
}

func (b *Balance) String() string {
	return fmt.Sprintf("%s => Profit: R$ %.2f, Debit: R$ %.2f, Accumulated: R$ %.2f", b.Date.Format("01/2006"), b.Profit, b.Debit, b.Accumulated)
}
