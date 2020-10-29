package main

import (
	"github.com/isongjosiah/lernen-api/api"
	"log"
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	//	start the api server
	a := &api.API{}

	log.Fatal(a.Serve())

}
