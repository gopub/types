package types

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
