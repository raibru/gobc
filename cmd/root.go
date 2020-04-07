package cmd

import (
	"github.com/raibru/gobaseconv/converter"
)

var bc converter.BaseContext

var prtPrefix bool
var prtVersion bool

func init() {
	rootCmd.PersistentFlags().StringVarP(&bc.Hex, "hex", "x", "", "convert string number from hexadecimal.\nRead from pipe use '.' as parameter")
	rootCmd.PersistentFlags().StringVarP(&bc.Dec, "dec", "d", "", "convert string number from decimal.\nRead from pipe use '.' as parameter")
	rootCmd.PersistentFlags().StringVarP(&bc.Oct, "oct", "o", "", "convert string number from octal.\nRead from pipe use '.' as parameter")
	rootCmd.PersistentFlags().StringVarP(&bc.Bin, "bin", "b", "", "convert string number from binary.\nRead from pipe use '.' as parameter")
	rootCmd.PersistentFlags().StringVarP(&bc.LeadZero, "lead-zero", "", "", "Print data with leading zeros\ndefined by a string and separated by ','\ne.g 'x:4,d:2,o:3,b:8'\n  hex: x:4 -> 4 leading zeros\n  dec: d:2 -> 2 leading zeros\n  oct: o:3 -> 3 leading zeros\n  bin: b:8 -> 8 leading zeros")
	rootCmd.PersistentFlags().BoolVarP(&prtPrefix, "prefix", "p", false, "display converted values with base prefix")
	rootCmd.PersistentFlags().BoolVarP(&prtVersion, "version", "v", false, "display gobc version")
}
