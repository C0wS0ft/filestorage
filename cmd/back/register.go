package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func handlerRegisterVolume(w http.ResponseWriter, r *http.Request) {
	v := new(Volume)
	err := json.NewDecoder(r.Body).Decode(&v)
	if err != nil {
		http.Error(w, "error parsing register request", http.StatusBadRequest)

		return
	}

	log.Printf("Registering volume %v", v.URL)
	volumes[v.URL] = struct{}{}

	_, _ = w.Write([]byte("ok"))
}
