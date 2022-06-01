package converter

import (
	"fmt"
	"strconv"
)

// BaseType hold base data
type BaseType struct {
	Number      uint64
	Value       string
	LeadZeros   string
	BaseId      string
	BaseNumber  int
	PrintFormat string
	IsInit      bool
}

// HexType hold data for hexadecimal types
type HexType struct {
	Base BaseType
}

// DecType hold data for decimal types
type DecType struct {
	Base BaseType
}

// OctType hold data for octal types
type OctType struct {
	Base BaseType
}

// BinType hold data for binary types
type BinType struct {
	Base BaseType
}

// Init HexType with default values
func (t *HexType) Init() {
	t.Base.BaseId = "x"
	t.Base.BaseNumber = 16
	t.Base.Value = "NaN"
	t.Base.LeadZeros = "0"
	t.Base.IsInit = true
}

// Init DecType with default values
func (t *DecType) Init() {
	t.Base.BaseId = "d"
	t.Base.BaseNumber = 10
	t.Base.Value = "NaN"
	t.Base.LeadZeros = "0"
	t.Base.IsInit = true
}

// Init OctType with default values
func (t *OctType) Init() {
	t.Base.BaseId = "o"
	t.Base.BaseNumber = 8
	t.Base.Value = "NaN"
	t.Base.LeadZeros = "0"
	t.Base.IsInit = true
}

// Init BinType with default values
func (t *BinType) Init() {
	t.Base.BaseId = "b"
	t.Base.BaseNumber = 2
	t.Base.Value = "NaN"
	t.Base.LeadZeros = "0"
	t.Base.IsInit = true
}

// SetupBy setup HexType by parameters from arguments
func (t *HexType) SetupBy(p *Parameters) error {
	e, _ := ParseLeadingZeros(p.LeadZeros, t.Base.BaseId)
	basePre := ""
	if p.UsePrefix {
		basePre = "0" + t.Base.BaseId
	}
	t.Base.PrintFormat = basePre + "%0" + e + "s"
	t.Base.Value = fmt.Sprintf(t.Base.PrintFormat, strconv.FormatUint(uint64(t.Base.Number), t.Base.BaseNumber))
	t.Base.LeadZeros = e
	return nil
}

// SetupBy setup DecType by parameters from arguments
func (t *DecType) SetupBy(p *Parameters) error {
	e, _ := ParseLeadingZeros(p.LeadZeros, t.Base.BaseId)
	basePre := ""
	if p.UsePrefix {
		basePre = "0" + t.Base.BaseId
	}
	t.Base.PrintFormat = basePre + "%0" + e + "s"
	t.Base.Value = fmt.Sprintf(t.Base.PrintFormat, strconv.FormatUint(uint64(t.Base.Number), t.Base.BaseNumber))
	t.Base.LeadZeros = e
	return nil
}

// SetupBy setup OctType by parameters from arguments
func (t *OctType) SetupBy(p *Parameters) error {
	e, _ := ParseLeadingZeros(p.LeadZeros, t.Base.BaseId)
	basePre := ""
	if p.UsePrefix {
		basePre = "0" + t.Base.BaseId
	}
	t.Base.PrintFormat = basePre + "%0" + e + "s"
	t.Base.Value = fmt.Sprintf(t.Base.PrintFormat, strconv.FormatUint(uint64(t.Base.Number), t.Base.BaseNumber))
	t.Base.LeadZeros = e
	return nil
}

// SetupBy setup BinType by parameters from arguments
func (t *BinType) SetupBy(p *Parameters) error {
	e, _ := ParseLeadingZeros(p.LeadZeros, t.Base.BaseId)
	basePre := ""
	if p.UsePrefix {
		basePre = "0" + t.Base.BaseId
	}
	t.Base.PrintFormat = basePre + "%0" + e + "s"
	t.Base.Value = fmt.Sprintf(t.Base.PrintFormat, strconv.FormatUint(uint64(t.Base.Number), t.Base.BaseNumber))
	t.Base.LeadZeros = e
	return nil
}

// String returns HexType base value as string
func (t *HexType) String() string {
	return t.Base.Value
}

// String returns DecType base value as string
func (t *DecType) String() string {
	return t.Base.Value
}

// String returns OctType base value as string
func (t *OctType) String() string {
	return t.Base.Value
}

// String returns BinType base value as string
func (t *BinType) String() string {
	return t.Base.Value
}
