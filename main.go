package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"strings"

	"g/src"
)

func getDataAndReturnResponse() {
	src.ParseJSON()
	// fmt.Println(src.All.Artists[0].Name)
	// return artist
}

func mainPage(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "404 Not Found page", 404)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "Method not Allowed", http.StatusMethodNotAllowed)
		return
	}
	temp, err := template.ParseFiles("./templates/index.html")
	if err != nil {
		http.Error(w, "Internal server error", 500)
		fmt.Println("Here")
		return
	}
	getDataAndReturnResponse()
	temp.Execute(w, src.All)
	return
}

func artistPage(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Method Not Allowed", 405)
		return
	}
	if !strings.HasPrefix(r.URL.Path, "/artist/") {
		http.NotFound(w, r)
		return
	}
	temp, err := template.ParseFiles("./templates/artist.html")
	if err != nil {
		http.Error(w, "Internal Server Error", 500)
		return
	}
	getDataAndReturnResponse()
	artistID, err := strconv.Atoi(r.URL.Path[8:])
	if err != nil {
		http.Error(w, "Bad Request", 400)
		log.Println(err)
		return
	}
	if artistID < 1 {
		http.Error(w, "Bad Request", 400)
		return
	} else if artistID > len(src.All.Artists) {
		// http.StatusBadRequest(w, r)
		http.Error(w, "Not Found", 404)
		return
	}
	artist := src.ArtistPageInfo(artistID - 1)
	temp.Execute(w, artist)
}

func main() {
	// getDataAndReturnResponse()
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("css/"))))
	http.HandleFunc("/", mainPage)
	http.HandleFunc("/artist/", artistPage)
	log.Fatal(http.ListenAndServe(":8080", nil))

}
