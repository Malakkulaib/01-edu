package main

import (
	"html/template"
	"net/http"
	"strconv"
)

func renderError(w http.ResponseWriter, status int) {
	w.WriteHeader(status)

	var page string
	switch status {
	case http.StatusNotFound:
		page = "templates/error404.html"
	case http.StatusBadRequest:
		page = "templates/error400.html"
	case http.StatusInternalServerError:
		page = "templates/error500.html"
	default:
		page = "templates/error500.html"
	}

	tmpl, err := template.ParseFiles(page)
	if err != nil {

		http.Error(w, "Internal Server Error: Error template missing", http.StatusInternalServerError)
		return
	}

	tmpl.Execute(w, nil)
}


func IndexHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		renderError(w, http.StatusNotFound)
		return
	}



	var artists []Artist
	err := fetchData("https://groupietrackers.herokuapp.com/api/artists", &artists)
	if err != nil {
		renderError(w, http.StatusInternalServerError)
		return
	}


	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		renderError(w, http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, artists)
}


func ArtistHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		renderError(w, http.StatusMethodNotAllowed)
		return
	}


	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil || id < 1 {
		renderError(w, http.StatusBadRequest)
		return
	}


	var artist Artist
	errArt := fetchData("https://groupietrackers.herokuapp.com/api/artists/"+idStr, &artist)
	if errArt != nil || artist.ID == 0 {
		renderError(w, http.StatusNotFound) 
		return
	}

	var rel Relation
	errRel := fetchData("https://groupietrackers.herokuapp.com/api/relation/"+idStr, &rel)
	if errRel != nil {
		renderError(w, http.StatusInternalServerError)
		return
	}


	data := struct {
		Artist   Artist
		Relation Relation
	}{
		Artist:   artist,
		Relation: rel,
	}

	tmpl, err := template.ParseFiles("templates/artist.html")
	if err != nil {
		renderError(w, http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, data)
}
