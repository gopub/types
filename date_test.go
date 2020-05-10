package types_test

import (
	"github.com/gopub/types"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDate_LastDayOfMonth(t *testing.T) {
	d := types.NewDate(2020, 1, 20)
	f := d.FirstDayOfMonth()
	assert.Equal(t, 2020, f.Year())
	assert.Equal(t, 1, f.Month())
	assert.Equal(t, 1, f.Day())
	l := d.LastDayOfMonth()
	assert.Equal(t, 2020, l.Year())
	assert.Equal(t, 1, l.Month())
	assert.Equal(t, 31, l.Day())
}
