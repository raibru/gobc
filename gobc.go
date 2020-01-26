package main

import (
	"os"

	"github.com/raibru/gobaseconv/cmd"
)

func main() {
	defer os.Exit(0)
	cmd.Execute()
}
