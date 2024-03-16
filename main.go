package main

import (
	"log"
	"net/http"

	"github.com/cepot-blip/image-processing-service/handlers"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/convert", handlers.ConvertImageHandler).Methods("POST")
	router.HandleFunc("/resize", handlers.ResizeImageHandler).Methods("POST")
	router.HandleFunc("/compress", handlers.CompressImageHandler).Methods("POST")

	serverAddr := ":8080"

	log.Printf("Server is starting and listening on %s...", serverAddr)
	log.Fatal(http.ListenAndServe(serverAddr, router))
}
