package types

import (
	"strconv"
)

// map[string]string
type MapStringString map[string]string

// get value for int
func (self MapStringString) Int(key string, def int) (int, error) {
	s, ok := self[key]
	if !ok {
		return def, nil
	}

	i, err := strconv.Atoi(s)
	if err != nil {
		return def, err
	}

	return i, nil
}
