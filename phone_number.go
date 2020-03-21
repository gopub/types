package types

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"strings"

	"github.com/gopub/conv"
	"github.com/gopub/gox/sql"
	"github.com/nyaruka/phonenumbers"
)

var _ driver.Valuer = (*PhoneNumber)(nil)
var _ sql.Scanner = (*PhoneNumber)(nil)

type PhoneNumber struct {
	Code      int    `json:"code"`
	Number    int64  `json:"number"`
	Extension string `json:"extension,omitempty"`
}

func (n *PhoneNumber) String() string {
	if len(n.Extension) == 0 {
		return fmt.Sprintf("+%d%d", n.Code, n.Number)
	}

	return fmt.Sprintf("+%d%d-%s", n.Code, n.Number, n.Extension)
}

func (n *PhoneNumber) InternationalFormat() string {
	pn, err := phonenumbers.Parse(n.String(), "")
	if err != nil {
		return ""
	}
	return phonenumbers.Format(pn, phonenumbers.INTERNATIONAL)
}

func (n *PhoneNumber) MaskString() string {
	nnBytes := []byte(fmt.Sprint(n.Number))
	maskLen := (len(nnBytes) + 2) / 3
	start := len(nnBytes) - 2*maskLen
	for i := 0; i < maskLen; i++ {
		nnBytes[start+i] = '*'
	}

	nn := string(nnBytes)

	if len(n.Extension) == 0 {
		return fmt.Sprintf("+%d%s", n.Code, nn)
	}

	return fmt.Sprintf("+%d%s-%s", n.Code, nn, n.Extension)
}

func (n *PhoneNumber) Scan(src interface{}) error {
	if src == nil {
		return nil
	}

	s, err := conv.ToString(src)
	if err != nil {
		return fmt.Errorf("parse string: %w", err)
	}
	if len(s) == 0 {
		return nil
	}

	fields, err := sql.ParseCompositeFields(s)
	if err != nil {
		return fmt.Errorf("parse composite fields %s: %w", s, err)
	}

	if len(fields) != 3 {
		return fmt.Errorf("parse composite fields %s: got %v", s, fields)
	}

	code, err := conv.ToInt(fields[0])
	if err != nil {
		return fmt.Errorf("parse code %s: %w", fields[0], err)
	}
	n.Code = int(code)
	n.Number, err = conv.ToInt64(fields[1])
	if err != nil {
		return fmt.Errorf("parse code %s: %w", fields[1], err)
	}
	n.Extension = fields[2]
	return nil
}

func (n *PhoneNumber) Value() (driver.Value, error) {
	if n == nil {
		return nil, nil
	}
	ext := strings.Replace(n.Extension, ",", "\\,", -1)
	s := fmt.Sprintf("(%d,%d,%s)", n.Code, n.Number, ext)
	return s, nil
}

func NewPhoneNumber(s string) (*PhoneNumber, error) {
	s = trimPhoneNumberString(s)
	pn, err := phonenumbers.Parse(s, "")
	if err != nil {
		return nil, err
	}

	if !phonenumbers.IsValidNumber(pn) {
		return nil, errors.New("invalid phone number")
	}

	return &PhoneNumber{
		Code:      int(pn.GetCountryCode()),
		Number:    int64(pn.GetNationalNumber()),
		Extension: pn.GetExtension(),
	}, nil
}

func NewPhoneNumberV2(s string, code int) (*PhoneNumber, error) {
	s = trimPhoneNumberString(s)
	pn, err := NewPhoneNumber(s)
	if err == nil {
		return pn, nil
	}
	if s[0] != '+' && code != 0 {
		s = fmt.Sprintf("+%d%s", code, s)
		return NewPhoneNumber(s)

	}
	return nil, err
}

func trimPhoneNumberString(s string) string {
	s = strings.Replace(s, "-", "", -1)
	s = strings.Replace(s, " ", "", -1)
	s = strings.Replace(s, "(", "", -1)
	s = strings.Replace(s, ")", "", -1)
	return s
}
