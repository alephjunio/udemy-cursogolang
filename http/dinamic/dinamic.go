package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func horaCerta(w http.ResponseWriter, r *http.Request) {
	s := time.Now().Format("02/01/2006 03:04:05")
	fmt.Fprintf(w, "<h1>A hora certa é: %s</h1>", s)
}

func main() {
	http.HandleFunc("/", horaCerta)
	log.Println("Listening on :3000")
	log.Println("Acesse: http://localhost:3000")
	log.Println("Executando.......")
	log.Fatal(http.ListenAndServe(":3000", nil))
}
