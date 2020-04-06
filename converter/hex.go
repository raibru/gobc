package converter

import (
	"bufio"
	"errors"
	"strconv"
)

//
// Use the Convertable interface
//

// HexConverter struct type for converting
type HexConverter struct{}

// ApplyPipeInput use data from pipe if assigned
func (c *HexConverter) ApplyPipeInput(s *bufio.Scanner, bc *BaseContext) error {
	if bc.Hex == "" {
		return nil
	}
	if bc.Hex == "." {
		for s.Scan() {
			bc.Hex = s.Text()
			return nil
		}
	}
	return errors.New("Hex parameter shall be set to '-' when read from pipe")
}

// ApplyContext apply an int64 value to binary value as string into base context bin data
func (c *HexConverter) ApplyContext(v int64, bc *BaseContext, m *map[string]string) BaseContext {
	bc.Hex = strconv.FormatInt(v, 16)
	e, ok := (*m)["x"]
	if !ok {
		e = "0"
	}
	bc.Hexlz = "%0" + e + "s"
	return *bc
}

// ParseInt64 parse the hexadecimal value into a int64 type
func (c *HexConverter) ParseInt64(bc *BaseContext) (int64, error) {
	val, err := strconv.ParseInt(bc.Hex, 16, 64)
	if err != nil {
		return -1, err
	}
	return val, nil
}

// Exists checks type specific BaseContext hexadecimal value
func (c *HexConverter) Exists(bc *BaseContext) bool {
	return len(bc.Hex) > 0
}
