package main

import (
	"log"
	"net/http"

	"golang-microservices/metadata/internal/controller/metadata"
	httphander "golang-microservices/metadata/internal/handler/http"
	"golang-microservices/metadata/internal/repository/memory"
)

func main() {
	log.Printf("Starting the movie metadata service")
	repo := memory.New()
	ctrl := metadata.New(repo)
	h := httphander.New(ctrl)
	http.Handle("/metadata", http.HandlerFunc(h.GetMetadata))
	if err := http.ListenAndServe(":8081", nil); err != nil {
		panic(err)
	}
}
