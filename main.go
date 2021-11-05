package main

import (
	"battlesnake-go-crazy/controllers"
	"battlesnake-go-crazy/engine"
	"log"
	"os"
	"strconv"
	"sync"
)

func main() {
	portEnv := os.Getenv("PORT")
	if len(portEnv) == 0 {
		portEnv = "8081"
	}

	port, err := strconv.Atoi(portEnv)
	if err != nil {
		log.Fatal("Invalid port given: ", portEnv)
	}

	var wg sync.WaitGroup
	randomHandlers := controllers.NewHandler(&engine.TotallyRandomEngine{})
	wg.Add(1)
	randomHandlers.StartListening(port, &wg)

	wg.Wait()
}
