package main

import (
	"crypto/sha256"
	"log"
	"net/http"
)

func getHashSum(d []byte) []byte {
	h := sha256.New()
	h.Write(d)

	return h.Sum(nil)
}

func main() {
	catalog = make(map[Name]*CatalogItem)
	volumes = make(map[URL]struct{})

	httpsMux := http.NewServeMux()
	httpsMux.HandleFunc("/v1/upload", handlerUpload)
	httpsMux.HandleFunc("/v1/download", handlerDownload)
	httpsMux.HandleFunc("/v1/register", handlerRegisterVolume)

	log.Println("Backend started")
	err := http.ListenAndServe(":8001", httpsMux)

	if err != nil {
		log.Printf("Error staring server: %s", err)
	}
}
