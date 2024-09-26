package main

type Dictionary map[string]string

const (
	ErrUnknownKey     = DictionaryErr("the key requested is not found")
	ErrDuplicateKey   = DictionaryErr("the key is already present")
	ErrFailedToUpdate = DictionaryErr("failed to update")
)

type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

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
		return ErrDuplicateKey
	default:
		return err
	}
}
