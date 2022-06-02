package cmd

import (
	"fmt"
	"io"
)

// var hold global variable
var (
	major     = "0"
	minor     = "10"
	patch     = "0"
	buildInfo = "-"
	appName   = "gobc - go base converter"
	author    = "raibru <github.com/raibru>"
	license   = "MIT License"
)

// PrintVersion prints the tool versions string
func PrintVersion(w io.Writer) {
	fmt.Fprintf(w, "%s - v%s.%s.%s\n", appName, major, minor, patch)
	fmt.Fprintf(w, "  Build : %s\n", buildInfo)
	fmt.Fprintf(w, "  Author : %s\n", author)
	fmt.Fprintf(w, "  License: %s\n\n", license)
}
