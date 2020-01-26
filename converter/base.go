package converter

import (
	"fmt"
	"io"
)

// BaseContext data structure
type BaseContext struct {
	Hex string
	Dec string
	Oct string
	Bin string
}

// PrintFormated prints base output in a simple format
func (bc *BaseContext) PrintFormated(w io.Writer) {
	fmt.Fprintf(w, "%s %s %s %s\n", bc.Hex, bc.Dec, bc.Oct, bc.Bin)
}

// PrintFormatedPrefix prints base output in a format with base prefix
func (bc *BaseContext) PrintFormatedPrefix(w io.Writer) {
	fmt.Fprintf(w, "0x%s 0d%s 0o%s 0b%s\n", bc.Hex, bc.Dec, bc.Oct, bc.Bin)
}
