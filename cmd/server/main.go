package main

import (
	"log"
	"net/http"
)

func main() {

	log.Println("Server up and running at http://127.0.0.1:9090")

	if err := http.ListenAndServe(":9090", http.FileServer(http.Dir("../../assets"))); err != nil {
		log.Fatal("Failed to start server", err)
	}

}
