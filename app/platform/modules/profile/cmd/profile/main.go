package main

import (
	"log"
	"net/http"
	"os"

	"github.com/hieunguyen1815/myself/modules/profile/internal/handler"
)

func main() {
	filepath := getEnv("RESUME_PATH", "./assets/default_profile.md")
	port := getEnv("PORT", "8080")

	mux := http.NewServeMux()
	handler.New(filepath).Register(mux)

	log.Printf("profile server listening on :%s (serving %s)", port, filepath)
	if err := http.ListenAndServe(":"+port, mux); err != nil {
		log.Fatal(err)
	}
}

func getEnv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}
