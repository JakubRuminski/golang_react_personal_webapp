package main

import (
	"log"
	"net/http"

	"github.com/jakubruminski/react_go/backend/go/public"
)

func main() {

	http.HandleFunc("/", public.HandleHomePage)

	log.Fatal(http.ListenAndServe(":9999", nil))
}
