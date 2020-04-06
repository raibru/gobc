package converter

import (
	"bufio"
	"errors"
	"strconv"
)

//
// Use the Convertable interface
//

// OctConverter struct type for converting
type OctConverter struct{}

// ApplyPipeInput use data from pipe if assigned
func (c *OctConverter) ApplyPipeInput(s *bufio.Scanner, bc *BaseContext) error {
	if bc.Oct == "" {
		return nil
	}
	if bc.Oct == "." {
		for s.Scan() {
			bc.Oct = s.Text()
			return nil
		}
	}
	return errors.New("Oct parameter shall be set to '-' when read from pipe")
}

// ApplyContext apply an int64 value to binary value as string into base context bin data
func (c *OctConverter) ApplyContext(v int64, bc *BaseContext, m *map[string]string) BaseContext {
	bc.Oct = strconv.FormatInt(v, 8)
	e, ok := (*m)["o"]
	if !ok {
		e = "0"
	}
	bc.Octlz = "%0" + e + "s"
	return *bc
}

// ParseInt64 parse the octal value into a int64 type
func (c *OctConverter) ParseInt64(bc *BaseContext) (int64, error) {
	val, err := strconv.ParseInt(bc.Oct, 8, 64)
	if err != nil {
		return -1, err
	}
	return val, nil
}

// Exists checks type specific BaseContext octal value
func (c *OctConverter) Exists(bc *BaseContext) bool {
	return len(bc.Oct) > 0
}
