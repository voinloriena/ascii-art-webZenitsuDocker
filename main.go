package main

import (
	"ascii-art-web/asciigo"
	"html/template"
	"io"
	"net/http"
	"os"
	"strings"
)

func getHello(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello, HTTP!\n")
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.html"))
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/static/", func(w http.ResponseWriter, r *http.Request) {
		filePath := "static" + strings.TrimPrefix(r.URL.Path, "/static/")

		// Check if the requested path is a directory
		if isDirectory(filePath) {
			http.Error(w, "Error:404 \nPage not found", http.StatusForbidden)
			return
		}

		// Serve the static file
		http.StripPrefix("/static/", http.FileServer(http.Dir("static"))).ServeHTTP(w, r)
	})

	mux.HandleFunc("/", indexHandler)
	mux.HandleFunc("/hello", getHello)
	mux.HandleFunc("/ascii-art", asciil)

	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		os.Exit(1)
	}
}

func isDirectory(path string) bool {
	info, err := os.Stat(path)
	return err == nil && info.IsDir()
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "Error:404 \nPage not Found", http.StatusNotFound)
		return
	}
	if r.Method != http.MethodGet {
		http.Error(w, "Error:405 \nMethod not allowed", http.StatusMethodNotAllowed)
		return
	}

	err := tpl.ExecuteTemplate(w, "index.html", nil)
	if err != nil {
		http.Error(w, "Error:500 \nInternal server error", http.StatusInternalServerError)
		return
	}
}

func asciil(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Redirect(w, r, "/", http.StatusSeeOther) // 303
		return
	}

	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Error:400 \nBad request: Unable to parse form", http.StatusBadRequest)
		return
	}

	input := r.FormValue("inpt")
	banner := r.FormValue("bnr")

	if input == "" || banner == "" {
		http.Error(w, "Error:400 \nInput and banner are required", http.StatusBadRequest)
		return
	}

	asciiOutput, err := asciigo.GenerateAsciiArt(input, banner)
	if err != nil {
		switch {
		case strings.Contains(err.Error(), "not found"):
			http.Error(w, "Error:404 \nNot found: Invalid banner selected", http.StatusNotFound)
		case strings.Contains(err.Error(), "not supported symbols"):
			http.Error(w, "Error:400 \nBad request: Invalid characters in input", http.StatusBadRequest)
		case strings.Contains(err.Error(), "bad request"):
			http.Error(w, "Error:400 \nBad request: Invalid banner selected", http.StatusBadRequest)
		case strings.Contains(err.Error(), "file is invalid"):
			http.Error(w, "Error:500 \nInternal server error: Banner file issue", http.StatusInternalServerError)
		case strings.Contains(err.Error(), "failed to open banner file"):
			http.Error(w, "Error:500 \nInternal server error: Banner file not found", http.StatusInternalServerError)
		default:
			http.Error(w, "Error:500 \nInternal server error", http.StatusInternalServerError)
		}
		return
	}

	asciiData := struct {
		Input  string
		Banner string
		Output string
	}{
		Input:  input,
		Banner: banner,
		Output: asciiOutput,
	}

	err = tpl.ExecuteTemplate(w, "ascii-art.html", asciiData)
	if err != nil {
		http.Error(w, "Error:500 \nInternal server error", http.StatusInternalServerError)
		return
	}
}
