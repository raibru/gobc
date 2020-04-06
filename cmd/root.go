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
	rootCmd.PersistentFlags().StringVarP(&bc.LeadZero, "lead-zero", "", "", "Print data with leading zero defined by string e.g 'x:4,d:2,o:3,b:8'\n  hex: x:4 - 4 leadings 0\n  dec: d:2 - 2 leadings 0\n  oct: o:3 - 3 leadings 0\n  bin: b:tttt8 - 8 leadings 0\nseparated by ','")
	rootCmd.PersistentFlags().BoolVarP(&prtPrefix, "prefix", "p", false, "display converted values with base prefix")
	rootCmd.PersistentFlags().BoolVarP(&prtVersion, "version", "v", false, "display gobc version")
}
