package cmd

import (
	"fmt"
	"io"
)

// PrintVersion prints the tool versions string
func PrintVersion(w io.Writer) {
	fmt.Fprintf(w, "gobc - go base converter - v0.9.4rc1\n")
	fmt.Fprintf(w, "  author : rbr <raibru@web.de>\n")
	fmt.Fprintf(w, "  license: MIT\n\n")
}
