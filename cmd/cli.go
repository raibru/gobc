package cmd

import (
	"fmt"
	"os"

	"github.com/raibru/gobaseconv/converter"
	"github.com/spf13/cobra"
)

var param converter.Parameters

var prtVersion bool

func init() {
	rootCmd.PersistentFlags().StringVarP(&param.Hex, "hex", "x", "", "convert string number from hexadecimal.\nRead from pipe use '.' as parameter")
	rootCmd.PersistentFlags().StringVarP(&param.Dec, "dec", "d", "", "convert string number from decimal.\nRead from pipe use '.' as parameter")
	rootCmd.PersistentFlags().StringVarP(&param.Oct, "oct", "o", "", "convert string number from octal.\nRead from pipe use '.' as parameter")
	rootCmd.PersistentFlags().StringVarP(&param.Bin, "bin", "b", "", "convert string number from binary.\nRead from pipe use '.' as parameter")
	rootCmd.PersistentFlags().StringVarP(&param.LeadZeros, "lead-zeros", "", "", "Print data with leading zeros\ndefined by a string and separated by ','\ne.g 'x:4,d:2,o:3,b:8'\n  hex: x:4 -> 4 leading zeros\n  dec: d:2 -> 2 leading zeros\n  oct: o:3 -> 3 leading zeros\n  bin: b:8 -> 8 leading zeros")
	rootCmd.PersistentFlags().BoolVarP(&param.UsePrefix, "prefix", "p", false, "display converted values with base prefix")
	rootCmd.PersistentFlags().BoolVarP(&prtVersion, "version", "v", false, "display gobc version")
}

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

	value, err := converter.ParseBaseValue(&param)
	if err != nil {
		return err
	}

	var bases = []converter.Converter{
		converter.NewConverterType[converter.HexType](value),
		converter.NewConverterType[converter.DecType](value),
		converter.NewConverterType[converter.OctType](value),
		converter.NewConverterType[converter.BinType](value),
	}

	//	fileInfo, _ := os.Stdin.Stat()
	//	if fileInfo.Mode()&os.ModeCharDevice == 0 {
	//		s := bufio.NewScanner(bufio.NewReader(os.Stdin))
	//		inerr := converter.ApplyPipeInput(s, &co, &bc)
	//		if inerr != nil {
	//			cmd.Help()
	//			return inerr
	//		}
	//	}

	for _, b := range bases {
		b.Init()
		err := b.SetupBy(&param)
		if err != nil {
			return err
		}
	}

	converter.PrintBases(bases, os.Stdout)

	return nil
}
