package main

import (
	"log"
	"net/http"
	"controllers/albums"
	"controllers/healthcheck"
)

func Routes() {
	log.Println("Routes loaded")

	Healthcheck()

	Albums()
}

func Healthcheck() {
	http.Handle("/", http.HandlerFunc(healthcheck.Status))
}

func Albums() {
	getAlbumsHandler := http.HandlerFunc(albums.GetAlbums)
	http.Handle("/albums", getAlbumsHandler)
}