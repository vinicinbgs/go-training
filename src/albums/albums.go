package albums

import (
	"io"
	"net/http"
	"log"
	"strconv"
)

type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

// postAlbums adds an album from JSON received in the request body.
// func postAlbums(c *gin.Context) {
//     var newAlbum album

//     // Call BindJSON to bind the received JSON to
//     // newAlbum.
//     if err := c.BindJSON(&newAlbum); err != nil {
//         return
//     }

//     // Add the new album to the slice.
//     albums = append(albums, newAlbum)
//     c.IndentedJSON(http.StatusCreated, newAlbum)
// }

func GetAlbums(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))

	if (err != nil) {
		log.Println(err)
	}

	log.Println(id)
	io.WriteString(w, albums[uint(id)].Title)
}
