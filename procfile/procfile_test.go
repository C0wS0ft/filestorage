package procfile

import (
	"crypto/sha256"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSplitJoinFile(t *testing.T) {
	d, err := os.ReadFile("/var/log/dmesg")
	require.NoError(t, err)

	h := sha256.New()
	h.Write(d)
	sum1 := h.Sum(nil)

	proc := NewFile(d, "dmesg", uint64(len(d)))
	splitted := proc.Split(6)

	joined := proc.Join(splitted)
	h.Reset()
	h.Write(joined)
	sum2 := h.Sum(nil)

	require.Equal(t, sum1, sum2)
}
