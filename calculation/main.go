package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/calculate", CalculateHandler)
	http.Handle("/", http.FileServer(http.Dir("pages/HTML")))

	fmt.Println("Serveren kjører på http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
