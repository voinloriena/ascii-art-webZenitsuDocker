package main

import (
	"ascii-art-web/asciigo"
	"html/template"
	"net/http"
	"os"
	"strings"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.html"))
}

func main() {
	mux := http.NewServeMux()

	// Обработка статики
	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	mux.Handle("/back/", http.StripPrefix("/back/", http.FileServer(http.Dir("back"))))

	// Основные маршруты
	mux.HandleFunc("/", indexHandler)
	mux.HandleFunc("/ascii-art", asciil)

	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		os.Exit(1)
	}
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		renderError(w, http.StatusNotFound, "Page not found")
		return
	}
	if r.Method != http.MethodGet {
		renderError(w, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}

	err := tpl.ExecuteTemplate(w, "index.html", nil)
	if err != nil {
		renderError(w, http.StatusInternalServerError, "Internal server error")
	}
}

func asciil(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	err := r.ParseForm()
	if err != nil {
		renderError(w, http.StatusBadRequest, "Bad request: Unable to parse form")
		return
	}

	input := r.FormValue("inpt")
	banner := r.FormValue("bnr")

	if input == "" || banner == "" {
		renderError(w, http.StatusBadRequest, "Input and banner are required")
		return
	}

	asciiOutput, err := asciigo.GenerateAsciiArt(input, banner)
	if err != nil {
		switch {
		case strings.Contains(err.Error(), "not found"):
			renderError(w, http.StatusNotFound, "Not found: Invalid banner selected")
		case strings.Contains(err.Error(), "not supported symbols"):
			renderError(w, http.StatusBadRequest, "Bad request: Invalid characters in input")
		case strings.Contains(err.Error(), "file is invalid"):
			renderError(w, http.StatusInternalServerError, "Internal server error: Banner file issue")
		default:
			renderError(w, http.StatusInternalServerError, "Internal server error")
		}
		return
	}

	data := struct {
		Input  string
		Banner string
		Output string
	}{
		Input:  input,
		Banner: banner,
		Output: asciiOutput,
	}

	err = tpl.ExecuteTemplate(w, "index.html", data)
	if err != nil {
		renderError(w, http.StatusInternalServerError, "Internal server error")
	}
}

func renderError(w http.ResponseWriter, code int, message string) {
	w.WriteHeader(code)
	data := struct {
		Code    int
		Message string
	}{
		Code:    code,
		Message: message,
	}
	err := tpl.ExecuteTemplate(w, "error.html", data)
	if err != nil {
		http.Error(w, message, code)
	}
}
