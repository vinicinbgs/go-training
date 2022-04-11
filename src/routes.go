package main

import (
	"io"
	"log"
	"net/http"
	"main/controllers"
)

func Routes() {
	log.Println("Routes loaded")

	http.Handle("/", http.HandlerFunc(Healthcheck))

	Albums()
}

func Healthcheck(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "OK")
}

func Albums() {
	getAlbumsHandler := http.HandlerFunc(albums.GetAlbums)
	
	http.Handle("/albums", getAlbumsHandler)
}