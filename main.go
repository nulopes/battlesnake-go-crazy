package main

import (
	"battlesnake-go-crazy/controllers"
	"battlesnake-go-crazy/engine"
	"log"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "8081"
	}

	handlers := controllers.NewHandler(&engine.TotallyRandomEngine{})

	http.HandleFunc("/", handlers.HandleIndex)
	http.HandleFunc("/start", handlers.HandleStart)
	http.HandleFunc("/move", handlers.HandleMove)
	http.HandleFunc("/end", handlers.HandleEnd)

	log.Printf("Starting Battlesnake Server at http://0.0.0.0:%s...\n", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
