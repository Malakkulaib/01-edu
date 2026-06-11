package main

import (
	"fmt"
	"net/http"
)

func main() {
	
	http.HandleFunc("/", IndexHandler)
	http.HandleFunc("/artist", ArtistHandler)
	
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	fi:= http.FileServer(http.Dir("./images"))
	http.Handle("/images/", http.StripPrefix("/images/", fi))



	fmt.Println("Server starting at http://localhost:1106")
	err := http.ListenAndServe(":1106", nil)
	if err != nil {
		fmt.Println("Server error:", err)
	}
}
