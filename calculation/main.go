package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/calculate", CalculateHandler)

	fmt.Println("Serveren kjører på http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
