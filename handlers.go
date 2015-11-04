package main

import (
	"html/template"
	"io"
	"net/http"
	"path"
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
	url := getQuery(r)
	if url == "" {
		http.Redirect(w, r, "/", 307)
		return
	}
	//userip := getUserIP(r)

	response, err := http.Get(url)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	defer response.Body.Close()

	_, err = io.Copy(w, response.Body) // read data from body and write to response writer

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

//func downloadFromUrl(url string) {
//fileName := getFileName(url)
//output, err := os.Create(fileName)
//if err != nil {
//return
//}
//defer output.Close()

//response, err := http.Get(url)
//if err != nil {
//return
//}
//defer response.Body.Close()

//n, err := io.Copy(output, response.Body)
//if err != nil {
//return
//}
//}

//func getFileName(url) {
//tokens := strings.Split(url, "/")
//fileName := tokens[len(tokens)-1]
//return fileName
//}

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
