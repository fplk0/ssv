package weekday

import (
	"bytes"
	"errors"
	"time"
)

type Consensus struct {
}

func (c *Consensus) ValidateValue(value []byte) error {
	day := time.Now().Weekday().String()

	// validate lambda
	if !bytes.Equal(value, []byte(day)) {
		return errors.New("msg value is wrong")
	}
	return nil
}
