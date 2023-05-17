package storage

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func TestMemStorage_Success(t *testing.T) {
	stor := NewMemStorage()

	data := "1234567890"
	id := "test"

	stor.Dump([]byte(data), id)
	rec, err := stor.Restore(id)
	require.NoError(t, err)
	require.Equal(t, data, string(rec))
}

func TestMemStorage_Error_NoSuchId(t *testing.T) {
	stor := NewMemStorage()

	_, err := stor.Restore(uuid.NewString())
	require.Error(t, err)
	require.Equal(t, err.Error(), "no such id")
}
