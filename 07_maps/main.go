package maps

import "errors"

var (
	ErrNotFound = errors.New("Word not in the dictionary")
	ErrExists   = errors.New("Already in the dictionary")
)

type Dictionary map[string]string

func (d Dictionary) Delete(key string) error {
	_, err := d.Search(key)

	switch err {
	case nil:
		delete(d, key)
	default:
		return err
	}

	return nil
}

func (d Dictionary) Update(key, value string) error {
	_, err := d.Search(key)

	switch err {
	case nil:
		d[key] = value
	default:
		return err
	}
	return nil
}

func (d Dictionary) Add(key, value string) error {
	_, err := d.Search(key)

	switch err {
	case ErrNotFound:
		d[key] = value
	case nil:
		return ErrExists
	default:
		return err
	}
	return nil
}

func (d Dictionary) Search(key string) (string, error) {
	definition, ok := d[key]

	if !ok {
		return "", ErrNotFound
	}

	return definition, nil
}
