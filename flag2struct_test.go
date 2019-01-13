package json2flag

import (
	"flag"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
	"time"
)

type CustomValue struct {
	v string
}

func (s CustomValue) String() string {
	return strings.ToLower(s.v)
}

func (s *CustomValue) Set(v string) error {
	s.v = strings.ToUpper(v)
	return nil
}

func TestFlag(t *testing.T) {
	type config struct {
		S  string
		B  bool
		I  int
		U  uint
		L  int64
		UL uint64
		F  float64
		D  time.Duration
		C  CustomValue
	}

	c := config{
		S:  "test",
		B:  true,
		I:  1,
		U:  2,
		L:  3,
		UL: 4,
		F:  5.67,
		D:  4 * time.Second,
		C:  CustomValue{v: "Custom"},
	}

	defaults := c

	Flag(&c, map[string]string{})
	flag.Parse()

	assert.Equal(t, defaults, c)

	assert.NoError(t, flag.Set("S", "passed"))
	assert.Equal(t, "passed", c.S)

	assert.NoError(t, flag.Set("B", "false"))
	assert.Equal(t, false, c.B)

	assert.NoError(t, flag.Set("I", "101"))
	assert.Equal(t, 101, c.I)

	assert.NoError(t, flag.Set("U", "102"))
	assert.Equal(t, uint(102), c.U)

	assert.NoError(t, flag.Set("L", "103"))
	assert.Equal(t, int64(103), c.L)

	assert.NoError(t, flag.Set("UL", "104"))
	assert.Equal(t, uint64(104), c.UL)

	assert.NoError(t, flag.Set("F", "105.67"))
	assert.Equal(t, 105.67, c.F)

	assert.NoError(t, flag.Set("D", "104s"))
	assert.Equal(t, 104*time.Second, c.D)

	assert.NoError(t, flag.Set("C", "cUsToM"))
	assert.Equal(t, CustomValue{v: "CUSTOM"}, c.C)
}

func TestFlagPrefixed(t *testing.T) {
	type config struct {
		S string
	}

	c := config{
		S: "test",
	}

	FlagPrefixed(&c, map[string]string{}, "m")
	flag.Parse()

	assert.NoError(t, flag.Set("m.S", "passed"))
	assert.Equal(t, "passed", c.S)
}
