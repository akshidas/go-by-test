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

func (d Dictionary) Add(key, value string) error {
	_, err := d.Search(key)

	switch err {
	case ErrUnknownKey:
		d[key] = value
		return nil
	case nil:
		return errors.New("duplicate key")
	default:
		return err
	}
}
