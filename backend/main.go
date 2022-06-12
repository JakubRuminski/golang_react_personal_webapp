package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/jakubruminski/react_go/backend/go/public"
)

var LOCATION_OF_STATIC string = "../ui/static"

func main() {

	http.HandleFunc("/", public.HandleHomePage)

	// Distributes css, js, img files when requested in html file.
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir(LOCATION_OF_STATIC))))

	info() // prints stuff to terminal

	log.Fatal(http.ListenAndServe(":9999", nil))
}

func info() {
	fmt.Print("\033[H\033[2J") // clears terminal window

	fmt.Println("")
	fmt.Println("##################################################")
	fmt.Println("Server running on localhost: http://localhost:9999")
	fmt.Println("##################################################")
}
