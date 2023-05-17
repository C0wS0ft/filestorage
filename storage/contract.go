package storage

type (
	Storage interface {
		Dump(d []byte, id string)
		Restore(id string) ([]byte, error)
	}

	memStorage struct {
		Data map[string][]byte
	}
)
