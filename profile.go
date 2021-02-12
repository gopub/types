package types

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/gopub/conv"
)

var (
	nickRegexp     = regexp.MustCompile("^[^ \n\r\t\f][^\n\r\t\f]{0,28}[^ \n\r\t\f]$")
	usernameRegexp = regexp.MustCompile("^[a-zA-Z][\\w\\.]{1,19}$")
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

func (g Gender) IsValid() bool {
	switch g {
	case Male, Female:
		return true
	default:
		return false
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
	Verified  bool   `json:"verified,omitempty"`
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
	Degree   EducationDegree `json:"degree,omitempty"`
	School   string          `json:"school,omitempty"`
	Major    string          `json:"major,omitempty"`
	StartAt  int64           `json:"started_at,omitempty"`
	EndAt    int64           `json:"end_at,omitempty"`
	Place    *Place          `json:"place,omitempty"`
	Verified bool            `json:"verified,omitempty"`
	Proofs   []string        `json:"documents,omitempty"`
}

type Work struct {
	Company  string   `json:"company,omitempty"`
	Title    string   `json:"title,omitempty"`
	Salary   int      `json:"salary,omitempty"`
	StartAt  int64    `json:"start_at,omitempty"`
	EndAt    int64    `json:"end_at,omitempty"`
	Place    *Place   `json:"place,omitempty"`
	Verified bool     `json:"verified,omitempty"`
	Proofs   []string `json:"documents,omitempty"`
}

type Account interface {
	AccountType() string
}

type Username string

func (u Username) AccountType() string {
	return "username"
}

func (u Username) IsValid() bool {
	return usernameRegexp.MatchString(string(u))
}
func (u Username) Normalize() Username {
	s := string(u)
	s = strings.TrimSpace(s)
	s = strings.ToLower(s)
	return Username(s)
}

type EmailAddress string

func (e EmailAddress) AccountType() string {
	return "email_address"
}

func (e EmailAddress) Normalize() EmailAddress {
	s := string(e)
	s = strings.TrimSpace(s)
	s = strings.ToLower(s)
	return EmailAddress(s)
}

func (e EmailAddress) IsValid() bool {
	return conv.IsEmailAddress(string(e))
}

type Nickname string

func (n Nickname) IsValid() bool {
	return nickRegexp.MatchString(string(n))
}

func (n Nickname) Normalize() Nickname {
	s := string(n)
	s = strings.TrimSpace(s)
	return Nickname(s)
}

func ParseAccount(s string) (Account, error) {
	s = strings.ToLower(s)
	s = strings.TrimSpace(s)
	if s == "" {
		return nil, errors.New("invalid account")
	}

	if pn, err := NewPhoneNumber(s); err == nil {
		return pn, nil
	}

	if conv.IsEmailAddress(s) {
		return EmailAddress(s), nil
	}

	if id, err := strconv.ParseInt(s, 10, 64); err == nil {
		return ID(id), nil
	}

	u := Username(s)
	if u.IsValid() {
		return u, nil
	}
	return nil, errors.New("invalid account")
}

func RandomEmailAddress() EmailAddress {
	s := RandomID().Short() + "@" + RandomID().Short() + ".com"
	return EmailAddress(s)
}

func RandomPhoneNumber() *PhoneNumber {
	return MustPhoneNumber(fmt.Sprintf("+861381234%04d", RandomID()%1e4))
}

func RandomNickname() Nickname {
	return Nickname(RandomID().Short())
}
