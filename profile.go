package types

import "fmt"

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

const (
	DriverLicense = "driver_license"
	NationalID    = "national_id"
	Passport      = "passport"
)

type PhotoID struct {
	Type      string `json:"type,omitempty"`
	Front     string `json:"front,omitempty"`
	Back      string `json:"back,omitempty"`
	Number    string `json:"number,omitempty"`
	IssuedAt  int64  `json:"issued_at,omitempty"`
	ExpiresAt int64  `json:"expires_at,omitempty"`
}

type EducationDegree int

const (
	PrimarySchool EducationDegree = 1 + iota
	JuniorSchool
	HighSchool
	JuniorCollege
	Bachelor
	Master
	PhD
)

type Education struct {
	Degree  EducationDegree `json:"degree,omitempty"`
	School  string          `json:"school,omitempty"`
	Major   string          `json:"major,omitempty"`
	StartAt int64           `json:"started_at,omitempty"`
	EndAt   int64           `json:"end_at,omitempty"`
}

type Work struct {
	Company string `json:"company"`
	Title   string `json:"title"`
	Salary  int    `json:"salary"`
	StartAt int64  `json:"start_at"`
	EndAt   int64  `json:"end_at"`
}
