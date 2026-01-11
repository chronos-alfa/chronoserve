package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

func main() {
	hostname := flag.String("host", "127.0.0.1", "Hostname to be used to serve the data")
	port := flag.Int("port", 8003, "Port to start listening on")
	folder := flag.String("folder", ".", "Where to start the server")

	flag.Parse()

	fullAddress := fmt.Sprintf("%s:%d", *hostname, *port)

	log.Printf("Starting to host the server at http://%s", fullAddress)

	handler := http.FileServer(http.Dir(*folder))

	log.Fatalln(http.ListenAndServe(fullAddress, handler))
}
