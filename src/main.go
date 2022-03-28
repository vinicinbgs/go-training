package main

import (
	"log"
	"net/http"
	"mypack/albums"
)

func main() {
	getAlbumsHandler := http.HandlerFunc(albums.GetAlbums)
	
	http.Handle("/albums", getAlbumsHandler)

	log.Println("Server Start")

	log.Fatal(http.ListenAndServe(":8080", nil))
}
