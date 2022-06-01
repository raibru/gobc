package converter

import (
	"errors"
	"fmt"
	"io"
	"regexp"
	"strconv"
	"strings"
)

// ConverterType interface defines an interface type constraint
type ConverterTypes interface {
	HexType | DecType | OctType | BinType
}

// NewConverterType create a new generic ConverterType object
func NewConverterType[T ConverterTypes](v uint64) *T {
	t := T{
		Base: BaseType{Number: v, IsInit: false},
	}
	return &t
}

// Converter interface handles converter behavior
type Converter interface {
	Init()
	SetupBy(*Parameters) error
	String() string
}

// Parameters hold data from arguments
type Parameters struct {
	Hex       string
	Dec       string
	Oct       string
	Bin       string
	LeadZeros string
	UsePrefix bool
}

//// ApplyPipeInput use data from pipe if assigned
//func ApplyPipeInput(s *bufio.Scanner, co *[]Convertable, bc *BaseContext) error {
//	for _, c := range *co {
//		perr := c.ApplyPipeInput(s, bc)
//		if perr != nil {
//			return perr
//		}
//	}
//	return nil
//}

// ParseBaseValue parse base number value from given parameter
func ParseBaseValue(p *Parameters) (uint64, error) {
	var once = []int{}
	var value uint64
	if p.Hex != "" {
		v, err := strconv.ParseUint(p.Hex, 16, 64)
		if err != nil {
			return 0, err
		}
		value = v
		once = append(once, 1)
	}

	if p.Dec != "" {
		v, err := strconv.ParseUint(p.Dec, 10, 64)
		if err != nil {
			return 0, err
		}
		value = v
		once = append(once, 1)
	}

	if p.Oct != "" {
		v, err := strconv.ParseUint(p.Oct, 8, 64)
		if err != nil {
			return 0, err
		}
		value = v
		once = append(once, 1)
	}

	if p.Bin != "" {
		v, err := strconv.ParseUint(p.Bin, 2, 64)
		if err != nil {
			return 0, err
		}
		value = v
		once = append(once, 1)
	}

	if len(once) == 0 {
		return 0, errors.New("need one parameter to convert into bases")
	} else if len(once) > 1 {
		return 0, errors.New("only one parameter shall be used")
	}

	return value, nil
}

// ParseLeadingZeros parse leading zeros for bases by given base id used by print format string
func ParseLeadingZeros(lzp string, bid string) (string, error) {
	m := make(map[string]string)
	if lzp != "" {
		elems := strings.Split(lzp, ",")
		r, err := regexp.Compile(`^[xdob]:\d+$`)
		if err != nil {
			return "0", err
		}

		for _, e := range elems {
			isMatch := r.MatchString(e)
			if !isMatch {
				return "0", errors.New("failure in print-format: " + lzp)
			}
			parts := strings.Split(e, ":")
			_, ok := m[parts[0]]
			if ok {
				return "0", errors.New("base type can be defined only once: " + lzp)
			}
			v := parts[1]
			m[parts[0]] = v
		}
	}

	e, ok := (m)[bid]
	if !ok {
		e = "0"
	}
	return e, nil
}

// PrintBases print all base string values to a writer
func PrintBases(bs []Converter, w io.Writer) {
	l := []string{}
	for _, b := range bs {
		l = append(l, b.String())
	}
	out := strings.Join(l, " ")
	fmt.Fprintf(w, "%s\n", out)
}
