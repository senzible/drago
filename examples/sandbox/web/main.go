package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/senzible/drago/host"
)

func main() {
	port := flag.String("p", "8080", "port to serve on")
	dir := flag.String("d", ".", "the directory of static file to host")
	flag.Parse()

	log.Printf("Serving UI %s on HTTP port: %s\n", *dir, *port)
	spaHandler := host.SpaHandler(*dir, "index.html")
	log.Fatal(http.ListenAndServe(":"+*port, spaHandler))
}
