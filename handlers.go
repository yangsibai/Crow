package main

import (
	"html/template"
	"net/http"
	"path"
	"strconv"
)

func Index(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles(path.Join("tmpl", "index.html"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := tmpl.Execute(w, nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func Download(w http.ResponseWriter, r *http.Request) {
	url := getQuery(url)
	if url == "" {
		http.Redirect(w, r, "/", 307)
		return
	}
	userip := getUserIP(r)

	downloadFromUrl(url)

	tmpl, err := template.ParseFiles(path.Join("tmpl", "download.html"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := tmpl.Execute(w, pageData); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func downloadFromUrl(url string) {
	fileName := getFileName(url)
	output, err := os.Create(fileName)
	if err != nil {
		return
	}
	defer output.Close()

	response, err := http.Get(url)
	if err != nil {
		return
	}
	defer response.Body.Close()

	n, err := io.Copy(output, response.Body)
	if err != nil {
		return
	}
}

func getFileName(url) {
	tokens := strings.Split(url, "/")
	fileName := tokens[len(tokens)-1]
	return fileName
}

func getUserIP(r *http.Request) string {
	if len(r.Header["X-Forwarded-For"]) != 0 {
		return r.Header["X-Forwarded-For"][0]
	} else {
		return r.RemoteAddr
	}
}

func getQuery(r *http.Request) string {
	words := r.URL.Query()["q"]
	return words[0]
}
