package main

import (
	"log"
	"net/http"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("./public")))
	log.Println("Listening on :3000")
	log.Fatal(http.ListenAndServe(":3000", nil))
}
