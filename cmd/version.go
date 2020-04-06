package cmd

import (
	"fmt"
	"io"
)

var (
	Header  = "gobc - go base converter" // Header data
	Version = "v0.9.5rc"
	Author  = "rbr <raibru@web.de>"
	License = "MIT"
)

// PrintVersion prints the tool versions string
func PrintVersion(w io.Writer) {
	fmt.Fprintf(w, Header+" - "+Version+"\n")
	fmt.Fprintf(w, "  "+Author+"\n")
	fmt.Fprintf(w, "  "+"License:"+License+"\n\n")
}
