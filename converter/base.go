package converter

import (
	"fmt"
	"io"
)

// BaseContext data structure
type BaseContext struct {
	Hex      string
	Hexlz    string
	Dec      string
	Declz    string
	Oct      string
	Octlz    string
	Bin      string
	Binlz    string
	LeadZero string
}

// PrintFormated prints base output in a simple format
func (bc *BaseContext) PrintFormated(w io.Writer) {
	pf := prepareFormat(bc, false)
	pf += "\n"
	fmt.Fprintf(w, pf, bc.Hex, bc.Dec, bc.Oct, bc.Bin)
}

// PrintFormatedPrefix prints base output in a format with base prefix
func (bc *BaseContext) PrintFormatedPrefix(w io.Writer) {
	pf := prepareFormat(bc, true)
	pf += "\n"
	fmt.Fprintf(w, pf, bc.Hex, bc.Dec, bc.Oct, bc.Bin)
}

func prepareFormat(bc *BaseContext, hasPrefix bool) string {
	f := ""
	if hasPrefix {
		f = "0x" + bc.Hexlz + " 0d" + bc.Declz + " 0o" + bc.Octlz + " 0b" + bc.Binlz
	} else {
		f = bc.Hexlz + " " + bc.Declz + " " + bc.Octlz + " " + bc.Binlz
	}
	return f
}
