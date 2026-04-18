package main

import (
	"log"
	"net/http"
	"os"

	"github.com/hieunguyen1815/myself/api"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("server listening on :%s", port)
	if err := http.ListenAndServe(":"+port, http.HandlerFunc(api.Handler)); err != nil {
		log.Fatal(err)
	}
}
