package main

import (
	"bytes"
	p "github.com/c0ws0ft/filestorage/procfile"
	"io"
	"log"
	"net/http"
	"sync"
)

func handlerDownload(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "no id set", http.StatusBadRequest)

		return
	}

	v, ok := catalog[id]
	if !ok {
		http.Error(w, "no such file", http.StatusNotFound)

		return
	}

	var pieces []*p.Piece

	var wg sync.WaitGroup
	var mtx sync.Mutex

	for _, v := range v.Volumes {
		wg.Add(1)
		go func(v Volume) {
			log.Printf("Getting part from %v", v.URL)
			resp, err := http.Get(v.URL + "/v1/download?id=" + id)
			if err != nil {
				http.Error(w, "error getting part from volume", http.StatusInternalServerError)

				return
			}
			body, err := io.ReadAll(resp.Body)
			if err != nil {
				http.Error(w, "error getting body", http.StatusInternalServerError)

				return
			}

			mtx.Lock()
			defer mtx.Unlock()
			pieces = append(pieces, &p.Piece{
				Data: body,
				Size: v.Size,
				Seq:  v.Seq,
			})
			wg.Done()
		}(v)
	}

	wg.Wait()
	res := p.New().Join(pieces)

	if !bytes.Equal(getHashSum(res), v.HashSum) {
		http.Error(w, "crc super error happened, sorry", http.StatusInternalServerError)

		return
	}

	_, _ = w.Write(res)
}
