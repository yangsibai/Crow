package main

import (
	"html/template"
	"io"
	"net/http"
	"net/url"
	"path"
	"strings"
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
	src := getQuery(r)
	if src == "" {
		http.Redirect(w, r, "/", 307)
		return
	}

	response, err := http.Get(src)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	defer response.Body.Close()

	w.Header().Set("Content-Length", response.Header().Get("Content-Length"))
	w.Header().Set("Content-Type", "application/octet-stream")
	w.Header().Set("Content-Disposition", "attachment;filename="+getFileName(src))

	_, err = io.Copy(w, response.Body) // read data from body and write to response writer

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func getFileName(src string) string {
	u, err := url.Parse(src)
	if err != nil {
		return "noname"
	}
	tokens := strings.Split(u.Path, "/")
	fileName := tokens[len(tokens)-1]
	if fileName == "" {
		return "noname"
	}
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
	words := r.URL.Query()["src"]
	return words[0]
}
