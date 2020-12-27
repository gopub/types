package types

import (
	"fmt"
)

// FullName defines user's full name
type FullName struct {
	First  string `json:"first"`
	Middle string `json:"middle"`
	Last   string `json:"last"`
}

func (n *FullName) String() string {
	if n == nil {
		return ""
	}
	return fmt.Sprintf("%s %s %s", n.First, n.Middle, n.Last)
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
