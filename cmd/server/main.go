package main

import (
	"log"
	"net/http"
)

func main() {

	log.Println("Server up and running at http://127.0.0.1:9090")

	log.Fatal(http.ListenAndServe(":9000", http.FileServer(http.Dir("../../assets"))))

}
