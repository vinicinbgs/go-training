package main

import (
	"log"
	"net/http"
)

func main() {
	Routes()

	log.Println("Server Start in http://localhost:8080")

	log.Fatal(http.ListenAndServe(":8080", nil))
}
