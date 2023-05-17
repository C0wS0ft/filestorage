package main

import (
	"bytes"
	p "github.com/c0ws0ft/filestorage/procfile"
	"log"
	"mime/multipart"
	"net/http"
	"sync"
)

func handlerUpload(w http.ResponseWriter, r *http.Request) {
	if len(volumes) == 0 {
		http.Error(w, "no storage volumes", http.StatusInsufficientStorage)

		return
	}

	if err := r.ParseMultipartForm(100 << 20); err != nil {
		http.Error(w, "error parsing multipart", http.StatusBadRequest)

		return
	}

	fil, head, _ := r.FormFile("file")
	fb := make([]byte, head.Size)
	_, err := fil.Read(fb)
	if err != nil {
		http.Error(w, "error processing multipart", http.StatusBadRequest)

		return
	}

	pieces := p.NewFile(fb, head.Filename, uint64(head.Size)).Split(uint64(len(volumes)))

	var item CatalogItem

	var wg sync.WaitGroup
	var mtx sync.Mutex
	i := 0

	for k := range volumes {
		wg.Add(1)
		go func(i uint64, k URL) {
			body := new(bytes.Buffer)
			writer := multipart.NewWriter(body)
			part, err := writer.CreateFormFile("file", head.Filename)
			if err != nil {
				log.Println("error creating multipart")

				return
			}
			_, _ = part.Write(pieces[i].Data)
			_ = writer.Close()

			log.Printf("Sending part to %v", k)
			_, err = http.Post(k+"/v1/upload", writer.FormDataContentType(), body)
			if err != nil {
				log.Println("Error sending post to volume")

				return
			}

			mtx.Lock()
			defer mtx.Unlock()
			item.Volumes = append(item.Volumes, Volume{
				URL:  k,
				Size: pieces[i].Size,
				Seq:  pieces[i].Seq,
			})
			wg.Done()
		}(uint64(i), k)
		i++
	}

	wg.Wait()

	item.Size = uint64(head.Size)
	item.HashSum = getHashSum(fb)
	catalog[head.Filename] = &item

	_, _ = w.Write([]byte("ok"))
}
