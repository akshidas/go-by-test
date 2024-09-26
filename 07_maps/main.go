package main

type Dictionary map[string]string

const (
	ErrUnknownKey   = DictionaryError("could not find the word you were looking for")
	ErrDuplicateKey = DictionaryError("could not add the key as duplicate exists")
)

type DictionaryError string

func (e DictionaryError) Error() string {
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
