package main

import (
	"flag"
	"log"
	"net/http"

	"golang.org/x/crypto/acme/autocert"
)

var domain = flag.String("domain", "dragffy.ro", "Your domain name")

func handler(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte("Hello / Welcome / Szia / Isten hozott / Bine ai venit!\n"))
}

func main() {
	flag.Parse()
	http.HandleFunc("/", handler)
	log.Fatal(http.Serve(autocert.NewListener(*domain), nil))
}
