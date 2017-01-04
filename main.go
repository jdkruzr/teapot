package main

import (
	// "fmt"
	"net/http"
	"log"
)

func handlerFunc(w http.ResponseWriter, r *http.Request) {
	log.Printf("%s %s %s", r.RemoteAddr, r.Method, r.RequestURI)
	w.WriteHeader(http.StatusTeapot)
}	
	
func main() {
	http.HandleFunc("/", handlerFunc)
	log.Println("Listening...")
	http.ListenAndServe(":3000", nil)
}