package main

import (
	"log"
	"net/http"
	"os"
)

const port = ":3000"

func main() {
	uiPath := os.Getenv("UI_PATH")
	if uiPath == "" {
		uiPath = "./ui/static"
	}

	fs := http.FileServer(http.Dir(uiPath))
	http.Handle("/", fs)

	log.Printf("Listening on %s ... [path=%s]\n", port, uiPath)
	log.Fatal(http.ListenAndServe(port, nil))
}
