package main

import (
	"flag"
	"log"
	"net/http"
)

var url, filesPath string

func init() {
	flag.StringVar(&url, "url", ":1234", "IP:Port address. ")
	flag.StringVar(&filesPath, "files", "./files", "Path to the files to serve them.") //todo improve this line to add ability set up few paths
}

func main() {
	flag.Parse()

	http.Handle("/", http.FileServer(http.Dir(filesPath)))

	log.Printf("Staring file server on '%s'", url)

	//serve files
	err := http.ListenAndServe(url, nil)
	if err != nil {
		log.Fatalf("Could not start file server by address: '%s'", url)
	}
}
