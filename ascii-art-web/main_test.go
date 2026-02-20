package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func performPost(t *testing.T, url, body string) *httptest.ResponseRecorder {
	req := httptest.NewRequest(http.MethodPost, url, strings.NewReader(body))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	rr := httptest.NewRecorder()
	asciiArtHandler(rr, req)
	return rr
}

func TestBanners(t *testing.T) {
	banners := []string{"standard", "shadow", "thinkertoy"}
	inputs := []string{"123", "<Hello> (World)!"}

	for _, banner := range banners {
		for _, text := range inputs {
			body := "text=" + text + "&banner=" + banner
			rr := performPost(t, "/ascii-art", body)

			if rr.Code != http.StatusOK {
				t.Errorf("Expected status 200, got %d for banner %s and input %s", rr.Code, banner, text)
			}
			if len(rr.Body.String()) == 0 {
				t.Errorf("Expected ASCII art output for banner %s and input %s", banner, text)
			}
		}
	}
}

func TestBadRequest(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/ascii-art", nil)
	rr := httptest.NewRecorder()
	asciiArtHandler(rr, req)

	if rr.Code != http.StatusBadRequest {
		t.Errorf("Expected 400 Bad Request, got %d", rr.Code)
	}
}

func TestNotFound(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/nonexistent-page", nil)
	rr := httptest.NewRecorder()
	http.DefaultServeMux.ServeHTTP(rr, req)

	if rr.Code != http.StatusNotFound {
		t.Errorf("Expected 404 Not Found, got %d", rr.Code)
	}
}

func TestInternalServerError(t *testing.T) {
	req := httptest.NewRequest(http.MethodPost, "/ascii-art", strings.NewReader("text=hello&banner=nonexist"))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	rr := httptest.NewRecorder()
	asciiArtHandler(rr, req)

	if rr.Code != http.StatusNotFound {
		t.Errorf("Expected 500 or 404 due to missing banner, got %d", rr.Code)
	}
}

func TestHomePage(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rr := httptest.NewRecorder()
	homeHandler(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("Expected 200 OK for home page, got %d", rr.Code)
	}
	body, _ := io.ReadAll(rr.Body)
	if !strings.Contains(string(body), "Enter your text") {
		t.Errorf("Home page does not contain textarea")
	}
}
