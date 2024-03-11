package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func uploadHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(2000)
	if err != nil {
		fmt.Fprintln(w, err)
		return
	}

	formdata := r.MultipartForm
	files := formdata.File["RemoteFile"]
	for i := range files {
		file, err := files[i].Open()
		defer file.Close()
		if err != nil {
			fmt.Fprintln(w, err)
			return
		}
		out, err := os.Create("UploadedImages/" + files[i].Filename)
		defer out.Close()
		if err != nil {
			fmt.Fprintf(w, "create file err!")
			return
		}
		_, err = io.Copy(out, file)
		if err != nil {
			fmt.Fprintln(w, err)
			return
		}
	}
}

func handFile(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		http.ServeFile(w, r, "./index.html")
	} else {
		http.ServeFile(w, r, "."+r.URL.Path)
	}
}

func main() {
	http.HandleFunc("/", handFile)
	http.HandleFunc("/upload", uploadHandler)
	http.ListenAndServe(":2024", nil)
}
