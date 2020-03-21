package types

import (
	"fmt"

	"github.com/shopspring/decimal"
)

// Money
type Money struct {
	Currency string          `json:"currency"`
	Amount   decimal.Decimal `json:"amount"`
}

func (m *Money) String() string {
	return fmt.Sprintf("%s %s", m.Currency, m.Amount.String())
}

func (m *Money) Equals(v *Money) bool {
	return m.Currency == v.Currency && m.Amount.Cmp(v.Amount) == 0
}
