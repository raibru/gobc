package converter

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"regexp"
	"strings"
)

// Convertable interface describe the convert behavior
type Convertable interface {
	ApplyPipeInput(s *bufio.Scanner, bc *BaseContext) error
	ApplyContext(v int64, ob *BaseContext, m *map[string]string) BaseContext
	ParseInt64(ob *BaseContext) (int64, error)
	Exists(ob *BaseContext) bool
}

// ApplyPipeInput use data from pipe if assigned
func ApplyPipeInput(s *bufio.Scanner, co *[]Convertable, bc *BaseContext) error {
	for _, c := range *co {
		perr := c.ApplyPipeInput(s, bc)
		if perr != nil {
			return perr
		}
	}
	return nil
}

// CreateConverter create a new convertable object depends of used parameter
func CreateConverter(co *[]Convertable, bc *BaseContext) (Convertable, error) {
	var si = []int{}
	for i, c := range *co {
		if c.Exists(bc) {
			si = append(si, i)
		}
	}
	if len(si) == 0 {
		return nil, errors.New("Need a parameter to convert into bases")
	} else if len(si) > 1 {
		return nil, errors.New("Only one parameter shall be used")
	}
	return (*co)[si[0]], nil
}

// ParseBaseValue parse the base value as int64 type value
func ParseBaseValue(c Convertable, bc *BaseContext) (int64, error) {
	return c.ParseInt64(bc)
}

// ApplyBaseContext apply convertable from int64 into base context
func ApplyBaseContext(cs *[]Convertable, v int64, bc *BaseContext) BaseContext {
	m, err := ParseLeadingZero(bc.LeadZero)
	if err != nil {
		fmt.Printf("Has format parsing error: %v", err)
		os.Exit(2) // change to error
	}
	for _, c := range *cs {
		c.ApplyContext(v, bc, m)
	}
	return *bc
}

// ParseLeadingZero evaluate the print format string and split it into a map
func ParseLeadingZero(pf string) (*map[string]string, error) {
	m := make(map[string]string)
	if pf != "" {
		elems := strings.Split(pf, ",")
		for _, e := range elems {
			isMatch, merr := regexp.MatchString("^[xdob]:\\d+$", e)
			if merr != nil {
				return &m, merr
			}
			if !isMatch {
				return &m, errors.New("Failure in print-format: " + pf)
			}
			parts := strings.Split(e, ":")
			_, ok := m[parts[0]]
			if ok {
				return &m, errors.New("Base type can be defined only once: " + pf)
			}
			v := parts[1]
			m[parts[0]] = v
		}
	}

	return &m, nil
}
