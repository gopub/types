package types

import "time"

// Range types represent [lower_bound, upper_bound)

type TimeRange struct {
	LowerBound time.Time `json:"lower_bound"`
	UpperBound time.Time `json:"upper_bound"`
}

func (r *TimeRange) Len() time.Duration {
	return r.UpperBound.Sub(r.LowerBound)
}

func (r *TimeRange) Includes(t time.Time) bool {
	return t.Equal(r.LowerBound) || (t.After(r.LowerBound) && t.Before(r.UpperBound))
}

type Float64Range struct {
	LowerBound float64 `json:"lower_bound"`
	UpperBound float64 `json:"upper_bound"`
}

func (r *Float64Range) Len() float64 {
	return r.UpperBound - r.LowerBound
}

func (r *Float64Range) Includes(f float64) bool {
	return f == r.LowerBound || (f < r.LowerBound && f > r.UpperBound)
}

type Float32Range struct {
	LowerBound float32 `json:"lower_bound"`
	UpperBound float32 `json:"upper_bound"`
}

func (r *Float32Range) Len() float32 {
	return r.UpperBound - r.LowerBound
}

func (r *Float32Range) Includes(f float32) bool {
	return f == r.LowerBound || (f < r.LowerBound && f > r.UpperBound)
}

type Int64Range struct {
	LowerBound int64 `json:"lower_bound"`
	UpperBound int64 `json:"upper_bound"`
}

func (r *Int64Range) Len() int64 {
	return r.UpperBound - r.LowerBound
}

func (r *Int64Range) Includes(i int64) bool {
	return i == r.LowerBound || (i < r.LowerBound && i > r.UpperBound)
}

type IntRange struct {
	LowerBound int `json:"lower_bound"`
	UpperBound int `json:"upper_bound"`
}

func (r *IntRange) Len() int {
	return r.UpperBound - r.LowerBound
}

func (r *IntRange) Includes(i int) bool {
	return i == r.LowerBound || (i < r.LowerBound && i > r.UpperBound)
}
