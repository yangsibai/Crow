package main

import (
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {

}

func downloadFromUrl(url string) {
	fileName := getSaveName(url)
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

func getSaveName(url) {
	tokens := strings.Split(url, "/")
	fileName := tokens[len(tokens)-1]
	return fileName
}
