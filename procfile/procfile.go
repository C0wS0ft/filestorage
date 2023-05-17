package procfile

import (
	"log"
	"sort"
)

func New() FileProcessor {
	return &myProcessor{}
}

func NewFile(d []byte, n string, s uint64) FileProcessor {
	return &myProcessor{
		Data: d,
		Name: n,
		Size: s,
	}
}

func (p *myProcessor) Split(n uint64) []*Piece {
	log.Printf("Splitting file [%v] to %v parts\n", p.Name, n)

	var pieceSize = uint64(p.Size / n)
	log.Println("Piece size: ", pieceSize)
	var left = p.Size - (pieceSize * n)
	lastSize := pieceSize + left
	log.Println("Last piece size: ", lastSize)

	var ret []*Piece

	for i := 0; i < int(n-1); i++ {
		var pc Piece
		pc.Data = p.Data[i*int(pieceSize) : (i+1)*int(pieceSize)]
		pc.Size = pieceSize
		pc.Seq = uint64(i)
		ret = append(ret, &pc)
	}

	var pc Piece
	pc.Data = p.Data[int(n-1)*int(pieceSize) : n*pieceSize+left]
	pc.Size = lastSize
	pc.Seq = n - 1
	ret = append(ret, &pc)

	return ret
}

func (p *myProcessor) Join(pcs []*Piece) []byte {
	sort.Slice(pcs, func(i, j int) bool {
		return pcs[i].Seq < pcs[j].Seq
	})

	ret := make([]byte, 0)

	for i := 0; i < len(pcs); i++ {
		ret = append(ret, pcs[i].Data...)
	}

	return ret
}
