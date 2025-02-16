package rw

import (
	"errors"
)

var ErrKeyNotFound = errors.New("key not found in the map")

type Data struct {
	M map[string]any
}

func (d *Data) ReadFromMap(key string) (string, error) {
	v, ok := d.M[key]
	if !ok {
		return "", ErrKeyNotFound
	}

	return v.(string), nil
}
