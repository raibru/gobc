package main

import (
	"os"

	"rbr.de/baseconv/cmd"
)

func main() {
	defer os.Exit(0)
	cmd.Execute()
}
