package storage

import "errors"

func NewMemStorage() Storage {
	return &memStorage{Data: make(map[string][]byte)}
}

func (m *memStorage) Dump(d []byte, id string) {
	m.Data[id] = d
}

func (m *memStorage) Restore(id string) ([]byte, error) {
	v, ok := m.Data[id]
	if !ok {
		return []byte{}, errors.New("no such id")
	}

	return v, nil
}
