package main

import (
	"fmt"
	"log"
	"net/http"
	"flux-router-engine/src/controllers"
)

func main() {
	fmt.Println("Flux Router Engine started")
	http.HandleFunc("/", controllers.Index)
	log.Fatal(http.ListenAndServe(":8080", nil))
}