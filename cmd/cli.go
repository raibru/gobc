package cmd

import (
	"bufio"
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"rbr.de/baseconv/converter"
)

var rootCmd = &cobra.Command{
	Use:   "gobc",
	Short: "Go Base Converter",
	Long:  `Exchange a given base-x number to hex, dec, oct, bin`,
	Run: func(cmd *cobra.Command, args []string) {
		if err := handleParam(cmd, args); err != nil {
			cmd.Help()
			fmt.Println("\nRoot command has parsing error: ", err)
			os.Exit(1)
		}
	},
}

// Execute the root command
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println("Execute root cmd has error", err)
		os.Exit(1)
	}
}

// handleParam parameter evaluation
func handleParam(cmd *cobra.Command, args []string) error {
	if prtVersion {
		PrintVersion(os.Stdout)
		return nil
	}

	var co = []converter.Convertable{}

	// ===============================================
	co = append(co, &converter.HexConverter{})
	co = append(co, &converter.DecConverter{})
	co = append(co, &converter.OctConverter{})
	co = append(co, &converter.BinConverter{})
	// ===============================================

	fileInfo, _ := os.Stdin.Stat()
	if fileInfo.Mode()&os.ModeCharDevice == 0 {
		s := bufio.NewScanner(bufio.NewReader(os.Stdin))
		inerr := converter.ApplyPipeInput(s, &co, &bc)
		if inerr != nil {
			cmd.Help()
			return inerr
		}
	}

	conv, cerr := converter.CreateConverter(&co, &bc)
	if cerr != nil {
		cmd.Help()
		return cerr
	}
	if conv == nil {
		cmd.Help()
		return nil
	}

	i, perr := converter.ParseBaseValue(conv, &bc)
	if perr != nil {
		return perr
	}

	out := converter.ApplyBaseContext(&co, i, &bc)

	if prtPrefix {
		out.PrintFormatedPrefix(os.Stdout)
	} else {
		out.PrintFormated(os.Stdout)
	}
	return nil
}
