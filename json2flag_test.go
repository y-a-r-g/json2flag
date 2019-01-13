package json2flag

import (
	"flag"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestReadConfigData(t *testing.T) {
	stringFlag := flag.String("stringFlag", "not works", "...")
	intFlag := flag.Int("intFlag", 0, "...")
	floatFlag := flag.Float64("floatFlag", 0, "...")
	boolFlag := flag.Bool("boolFlag", false, "...")
	durationFlag := flag.Duration("durationFlag", 0, "...")
	defaultFlag := flag.String("defaultFlag", "default", "...")

	dataString := "{\"stringFlag\": \"works\", \"intFlag\": 42, \"floatFlag\": 3.1415, \"boolFlag\": true, \"durationFlag\": \"5s\"}"

	flag.Parse()
	assert.NoError(t, ReadConfigString(dataString))

	assert.Equal(t, "works", *stringFlag)
	assert.Equal(t, 42, *intFlag)
	assert.Equal(t, 3.1415, *floatFlag)
	assert.Equal(t, true, *boolFlag)
	assert.Equal(t, 5*time.Second, *durationFlag)
	assert.Equal(t, "default", *defaultFlag)
}

func TestNestedConfig(t *testing.T) {
	stringFlag := flag.String("a.b", "not works", "...")
	dataString := "{\"a\": {\"b\": \"passed\"}}"

	flag.Parse()
	assert.NoError(t, ReadConfigString(dataString))
	assert.Equal(t, "passed", *stringFlag)
}
