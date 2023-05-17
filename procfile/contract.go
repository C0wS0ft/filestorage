package procfile

type (
	FileProcessor interface {
		Split(n uint64) []*Piece
		Join([]*Piece) []byte
	}

	myProcessor struct {
		Data []byte
		Name string
		Size uint64
	}

	Piece struct {
		Data []byte
		Size uint64
		Seq  uint64
	}
)
