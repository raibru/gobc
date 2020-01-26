package converter

import (
	"bufio"
	"errors"
	"strconv"
)

//
// Use the Convertable interface
//

// DecConverter struct type for converting
type DecConverter struct{}

// ApplyPipeInput use data from pipe if assigned
func (c *DecConverter) ApplyPipeInput(s *bufio.Scanner, bc *BaseContext) error {
	if bc.Dec == "" {
		return nil
	}
	if bc.Dec == "." {
		for s.Scan() {
			bc.Dec = s.Text()
			return nil
		}
	}
	return errors.New("Dec parameter shall be set to '-' when read from pipe")
}

// ApplyContext apply an int64 value to decimal value as string into base context dec data
func (c *DecConverter) ApplyContext(v int64, bc *BaseContext) BaseContext {
	bc.Dec = strconv.FormatInt(v, 10)
	return *bc
}

// ParseInt64 parse the decimal value into a int64 type
func (c *DecConverter) ParseInt64(bc *BaseContext) (int64, error) {
	val, err := strconv.ParseInt(bc.Dec, 10, 64)
	if err != nil {
		return -1, err
	}
	return val, nil
}

// Exists checks type specific OutBae decimal value
func (c *DecConverter) Exists(bc *BaseContext) bool {
	return len(bc.Dec) > 0
}
