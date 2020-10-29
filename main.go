package main

import (
	"github.com/isongjosiah/lernen-api/api"
	"log"
)

func main() {
	//	start the api server
	a := &api.API{}

	log.Fatal(a.Serve())

}
