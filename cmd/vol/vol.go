package main

import (
	"log"
	"net/http"
	"os"

	s "github.com/c0ws0ft/filestorage/storage"
)

func handlerUpload(w http.ResponseWriter, r *http.Request) {
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

	log.Printf("Dumping %v, size %v", head.Filename, head.Size)
	storage.Dump(fb, head.Filename)

	_, _ = w.Write([]byte("ok"))
}

func handlerDownload(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		log.Println("Not a GET")
		http.Error(w, "404 not found.", http.StatusBadRequest)

		return
	}

	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "no id set", http.StatusBadRequest)

		return
	}

	data, err := storage.Restore(id)
	if err != nil {
		http.Error(w, "id not found", http.StatusNotFound)

		return
	}

	_, _ = w.Write(data)
}

func main() {
	if len(os.Args) != 2 {
		log.Println("Usage: program <:8002>")

		return
	}

	storage = s.NewMemStorage()

	httpsMux := http.NewServeMux()
	httpsMux.HandleFunc("/v1/upload", handlerUpload)
	httpsMux.HandleFunc("/v1/download", handlerDownload)

	log.Println("Volume started at", os.Args[1])
	err := http.ListenAndServe(os.Args[1], httpsMux)

	if err != nil {
		log.Printf("Error staring volume: %v", err.Error())
	}
}
