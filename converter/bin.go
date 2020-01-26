package converter

import (
	"bufio"
	"errors"
	"strconv"
)

//
// Use the Convertable interface
//

// BinConverter struct type for converting
type BinConverter struct{}

// ApplyPipeInput use data from pipe if assigned
func (c *BinConverter) ApplyPipeInput(s *bufio.Scanner, bc *BaseContext) error {
	if bc.Bin == "" {
		return nil
	}
	if bc.Bin == "." {
		for s.Scan() {
			bc.Bin = s.Text()
			return nil
		}
	}
	return errors.New("Bin parameter shall be set to '-' when read from pipe")
}

// ApplyContext apply an int64 value to binary value as string into base context bin data
func (c *BinConverter) ApplyContext(v int64, bc *BaseContext) BaseContext {
	bc.Bin = strconv.FormatInt(v, 2)
	return *bc
}

// ParseInt64 parse the binary value into a int64 type
func (c *BinConverter) ParseInt64(bc *BaseContext) (int64, error) {
	val, err := strconv.ParseInt(bc.Bin, 2, 64)
	if err != nil {
		return -1, err
	}
	return val, nil
}

// Exists checks type specific BaseContext binary value
func (c *BinConverter) Exists(bc *BaseContext) bool {
	return len(bc.Bin) > 0
}
