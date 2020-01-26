package cmd

import (
	"rbr.de/baseconv/converter"
)

var bc converter.BaseContext

var prtPrefix bool
var prtVersion bool

func init() {
	rootCmd.PersistentFlags().StringVarP(&bc.Hex, "hex", "x", "", "convert string number from hexadecimal.\nRead from pipe use '.' as parameter")
	rootCmd.PersistentFlags().StringVarP(&bc.Dec, "dec", "d", "", "convert string number from decimal.\nRead from pipe use '.' as parameter")
	rootCmd.PersistentFlags().StringVarP(&bc.Oct, "oct", "o", "", "convert string number from octal.\nRead from pipe use '.' as parameter")
	rootCmd.PersistentFlags().StringVarP(&bc.Bin, "bin", "b", "", "convert string number from binary.\nRead from pipe use '.' as parameter")
	rootCmd.PersistentFlags().BoolVarP(&prtPrefix, "prefix", "p", false, "display converted values with base prefix")
	rootCmd.PersistentFlags().BoolVarP(&prtVersion, "version", "v", false, "display gobc version")
}
