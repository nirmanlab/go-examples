package main

import (
	"fmt"
	"log"
	"net/http"
)

type Article struct {
	Title       string `json:name`
	Description string `json:desc`
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Homepage Endpoint Hit")
}
func handelRequest() {
	http.HandleFunc("/", homePage)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
func main() {
	handelRequest()
}
