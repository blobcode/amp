package main

import (
	"embed"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

// embedded files
//go:embed content/*
var f embed.FS

const port = 8080 // port to run on

func main() {

	// use enviroment variable for port if exists
	envport, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		envport = port
	}

	// start
	log.Println("welcome to amp")
	log.Println("listening on http://localhost:" + strconv.Itoa(envport))

	// register handler
	http.HandleFunc("/", handler)

	err = http.ListenAndServe("localhost:"+strconv.Itoa(envport), nil)

	if err != nil {
		log.Fatal(err)
	}
}

func handler(w http.ResponseWriter, r *http.Request) {

	// index hack
	if r.URL.Path == "/" {
		r.URL.Path = "index.html"
	}
	fp := "content/" + strings.TrimPrefix(filepath.Clean(r.URL.Path), `/`)

	p, err := f.ReadFile(fp)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	w.Write(p)

	log.Println("served", fp)
}
