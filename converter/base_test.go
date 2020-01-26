package converter

import (
	"strings"
	"testing"
)

func Test_BasePrintFormat(t *testing.T) {
	//t.Fatal("Check Failure")
	// Given
	tables := []struct {
		input  string
		expect string
	}{
		{"0", "0 0 0 0"},
		{"1", "1 1 1 1"},
		{"2", "2 2 2 10"},
		{"7", "7 7 7 111"},
		{"10", "a 10 12 1010"},
		{"16", "10 16 20 10000"},
	}

	for _, table := range tables {

		// When
		bc := BaseContext{
			Dec: table.input}

		var co = []Convertable{}
		co = append(co, &DecConverter{})
		co = append(co, &HexConverter{})
		co = append(co, &OctConverter{})
		co = append(co, &BinConverter{})

		conv, err := CreateConverter(&co, &bc)

		if err != nil || conv == nil {
			t.Fatal("Fatal: can not create converter")
		}

		i, perr := ParseBaseValue(conv, &bc)

		if perr != nil {
			t.Fatalf("Fatal: can not parse base value '%v'", table.input)
		}

		result := ApplyBaseContext(&co, i, &bc)

		b := &strings.Builder{}
		b.Reset()
		result.PrintFormated(b)
		output := strings.TrimSpace(b.String())

		// Then
		if table.expect != output {
			t.Errorf("Failure: PrintFormat expect '%s' but get '%v'", table.expect, output)
		}
	}
}

func Test_BasePrintFormatPrefix(t *testing.T) {
	//t.Fatal("Check Failure")
	// Given
	tables := []struct {
		input  string
		expect string
	}{
		{"0", "0x0 0d0 0o0 0b0"},
		{"1", "0x1 0d1 0o1 0b1"},
		{"2", "0x2 0d2 0o2 0b10"},
		{"7", "0x7 0d7 0o7 0b111"},
		{"10", "0xa 0d10 0o12 0b1010"},
		{"16", "0x10 0d16 0o20 0b10000"},
	}

	for _, table := range tables {

		// When
		bc := BaseContext{
			Dec: table.input}

		var co = []Convertable{}
		co = append(co, &DecConverter{})
		co = append(co, &HexConverter{})
		co = append(co, &OctConverter{})
		co = append(co, &BinConverter{})

		conv, err := CreateConverter(&co, &bc)

		if err != nil || conv == nil {
			t.Fatal("Fatal: can not create converter")
		}

		i, perr := ParseBaseValue(conv, &bc)

		if perr != nil {
			t.Fatalf("Fatal: can not parse base value '%v'", table.input)
		}

		result := ApplyBaseContext(&co, i, &bc)

		b := &strings.Builder{}
		b.Reset()
		result.PrintFormatedPrefix(b)
		output := strings.TrimSpace(b.String())

		// Then
		if table.expect != output {
			t.Errorf("Failure: PrintFormat expect '%s' but get '%v'", table.expect, output)
		}
	}
}
