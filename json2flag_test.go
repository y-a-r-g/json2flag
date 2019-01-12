package json2flag

import (
	"flag"
	"testing"
)

func TestReadConfigData(t *testing.T) {
	stringFlag := flag.String("stringFlag", "not works", "...")
	intFlag := flag.Int("intFlag", 0, "...")
	floatFlag := flag.Float64("floatFlag", 0, "...")
	boolFlag := flag.Bool("boolFlag", false, "...")

	flag.Parse()
	if err := ReadConfigString("{\"stringFlag\": \"works\", \"intFlag\": 42, \"floatFlag\": 3.1415, \"boolFlag\": true}"); err != nil {
		panic(err)
	}

	if *stringFlag != "works" || *intFlag != 42 || *floatFlag != 3.1415 || !*boolFlag {
		t.Error("config was not read")
	}
}
