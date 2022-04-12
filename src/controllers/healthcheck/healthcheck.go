package healthcheck

import (
	"io"
	"log"
	"net/http"
	"encoding/json"
)

type S_Healthcheck struct {
	Status int
	Message string
}

func Status(w http.ResponseWriter, r *http.Request) {
	log.Println("OK")
	object := S_Healthcheck{200, "OK"}
	response, _ := json.Marshal(object)
	io.WriteString(w, string(response))
}