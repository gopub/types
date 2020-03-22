package types

import (
	"fmt"
)

// FullName defines user's full name
type FullName struct {
	FirstName  string `json:"first_name"`
	MiddleName string `json:"middle_name"`
	LastName   string `json:"last_name"`
}

func (n *FullName) String() string {
	return fmt.Sprintf("%s %s %s", n.FirstName, n.MiddleName, n.LastName)
}

type Gender int

const (
	Male   Gender = 1
	Female Gender = 2
)

func (g Gender) String() string {
	switch g {
	case Male:
		return "male"
	case Female:
		return "female"
	default:
		return "unknown"
	}
}
