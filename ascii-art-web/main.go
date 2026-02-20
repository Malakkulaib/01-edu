package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/home", homeHandler)
	http.HandleFunc("/about", aboutHandler)
	http.HandleFunc("/test", testHandler)
	http.HandleFunc("/ascii-art", asciiArtHandler)


	fs := http.FileServer(http.Dir("./images"))
	http.Handle("/images/", http.StripPrefix("/images/", fs))


	fmt.Println("Server running on http://localhost:3112")
	err := http.ListenAndServe(":3112", nil)
	if err != nil {
		fmt.Println("Server error:", err)
	}
}
