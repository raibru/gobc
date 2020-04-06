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

// ApplyContext apply an int64 value to binary value as string into base context bin data
func (c *DecConverter) ApplyContext(v int64, bc *BaseContext, m *map[string]string) BaseContext {
	bc.Dec = strconv.FormatInt(v, 10)
	e, ok := (*m)["d"]
	if !ok {
		e = "0"
	}
	bc.Declz = "%0" + e + "s"
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
