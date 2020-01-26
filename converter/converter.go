package converter

import (
	"bufio"
	"errors"
)

// Convertable interface describe the convert behavior
type Convertable interface {
	ApplyPipeInput(s *bufio.Scanner, bc *BaseContext) error
	ApplyContext(v int64, ob *BaseContext) BaseContext
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
	for _, c := range *cs {
		c.ApplyContext(v, bc)
	}
	return *bc
}
