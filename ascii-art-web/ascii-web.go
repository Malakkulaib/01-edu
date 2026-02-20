package main

import (
	"html/template"
	"net/http"
	"strings"
)

func renderError(w http.ResponseWriter, status int) {
	w.WriteHeader(status)

	var page string

	switch status {
	case http.StatusNotFound:
		page = "error404.html"
	case http.StatusBadRequest:
		page = "error400.html"
	case http.StatusInternalServerError:
		page = "error500.html"
	default:
		page = "error500.html"
	}

	tmpl, err := template.ParseFiles(page)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	tmpl.Execute(w, nil)
}

func renderTemplate(w http.ResponseWriter, filename string, data any) {
	tmpl, err := template.ParseFiles(filename)
	if err != nil {
		renderError(w, http.StatusNotFound)
		return
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		renderError(w, http.StatusInternalServerError)
		return
	}
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		renderError(w, http.StatusNotFound)
		return
	}
	renderTemplate(w, "index.html", nil)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "home.html", nil)
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "about.html", nil)
}

func testHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "test.html", nil)
}

func asciiArtHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		renderError(w, http.StatusBadRequest)
		return
	}

	text := r.FormValue("text")
	if text == "" {
		renderError(w, http.StatusBadRequest)
		return
	}

	text = strings.ReplaceAll(text, `\n`, "\n")

	bannerName := r.FormValue("banner")
	bannerFile := bannerName + ".txt"

	banner, err := loadBanner(bannerFile)
	if err != nil {
		renderError(w, http.StatusNotFound)
		return
	}

	lines := strings.Split(text, "\n")

	var result strings.Builder
	result.WriteString("\n")

	for _, line := range lines {
		if line == "" {
			continue
		}
		result.WriteString(renderASCII(line, banner))
		result.WriteString("\n")
	}

	renderTemplate(w, "test.html", map[string]string{
		"Result": result.String(),
	})
}
