package main

import "errors"

type Dictionary map[string]string

var ErrUnknownKey = errors.New("unknown key")

func (d Dictionary) Search(key string) (string, error) {
	value, ok := d[key]
	if ok {
		return value, nil
	}

	return "", ErrUnknownKey
}
